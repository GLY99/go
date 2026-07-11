package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	masterpb "mygrpc/proto"
	pb "mygrpc/proto"
)

// AgentInfo 存储 Agent 信息
type AgentInfo struct {
	AgentID       string
	AgentName     string
	Address       string
	MaxTasks      int32
	Capabilities  []string
	RunningTasks  int32
	Load          int32
	Status        string // "healthy", "busy", "unhealthy"
	LastHeartbeat time.Time
	Conn          *grpc.ClientConn
	Client        pb.AgentServiceClient
	Mutex         sync.RWMutex
}

// MasterServer 实现 Master 服务
type MasterServer struct {
	masterpb.UnimplementedMasterServiceServer
	agents      map[string]*AgentInfo
	agentMutex  sync.RWMutex
	taskCounter int32
}

func NewMasterServer() *MasterServer {
	return &MasterServer{
		agents: make(map[string]*AgentInfo),
	}
}

// RegisterAgent Agent 注册
func (s *MasterServer) RegisterAgent(ctx context.Context, req *masterpb.RegisterRequest) (*masterpb.RegisterResponse, error) {
	s.agentMutex.Lock()
	defer s.agentMutex.Unlock()

	// 检查是否已注册
	if _, exists := s.agents[req.AgentId]; exists {
		return &masterpb.RegisterResponse{
			Success: false,
			Message: "Agent 已注册",
		}, nil
	}

	// 连接到 Agent
	conn, err := grpc.Dial(req.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &masterpb.RegisterResponse{
			Success: false,
			Message: fmt.Sprintf("连接 Agent 失败: %v", err),
		}, nil
	}

	agent := &AgentInfo{
		AgentID:       req.AgentId,
		AgentName:     req.AgentName,
		Address:       req.Address,
		MaxTasks:      req.MaxTasks,
		Capabilities:  req.Capabilities,
		Status:        "healthy",
		LastHeartbeat: time.Now(),
		Conn:          conn,
		Client:        pb.NewAgentServiceClient(conn),
	}

	s.agents[req.AgentId] = agent

	log.Printf("[Master] ✅ Agent 注册成功: %s (%s)", req.AgentName, req.AgentId)
	return &masterpb.RegisterResponse{
		Success: true,
		Message: "注册成功",
	}, nil
}

// Heartbeat 处理心跳
func (s *MasterServer) Heartbeat(ctx context.Context, req *masterpb.HeartbeatRequest) (*masterpb.HeartbeatResponse, error) {
	s.agentMutex.RLock()
	agent, exists := s.agents[req.AgentId]
	s.agentMutex.RUnlock()

	if !exists {
		return &masterpb.HeartbeatResponse{
			Success: false,
			Action:  "stop",
		}, nil
	}

	agent.Mutex.Lock()
	agent.RunningTasks = req.RunningTasks
	agent.Load = req.Load
	agent.LastHeartbeat = time.Now()
	agent.Status = "healthy"
	if req.RunningTasks >= agent.MaxTasks {
		agent.Status = "busy"
	}
	agent.Mutex.Unlock()

	return &masterpb.HeartbeatResponse{
		Success: true,
		Action:  "continue",
	}, nil
}

// SelectAgent 选择最优 Agent (负载均衡)
func (s *MasterServer) SelectAgent(capability string) (*AgentInfo, error) {
	s.agentMutex.RLock()
	defer s.agentMutex.RUnlock()

	var selected *AgentInfo
	var minLoad int32 = 1000

	for _, agent := range s.agents {
		agent.Mutex.RLock()
		// 检查 Agent 是否健康
		if agent.Status == "unhealthy" {
			agent.Mutex.RUnlock()
			continue
		}

		// 检查能力
		hasCapability := false
		for _, cap := range agent.Capabilities {
			if cap == capability || capability == "" {
				hasCapability = true
				break
			}
		}

		// 检查是否还有空余容量
		canAccept := agent.RunningTasks < agent.MaxTasks

		agent.Mutex.RUnlock()

		if hasCapability && canAccept {
			agent.Mutex.RLock()
			currentLoad := agent.Load
			agent.Mutex.RUnlock()

			if currentLoad < minLoad {
				minLoad = currentLoad
				selected = agent
			}
		}
	}

	if selected == nil {
		return nil, fmt.Errorf("没有可用的 Agent")
	}

	return selected, nil
}

// ExecuteTaskOnAgent 在指定 Agent 上执行任务
func (s *MasterServer) ExecuteTaskOnAgent(agent *AgentInfo, taskID, taskType, input string, params map[string]string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := &pb.TaskRequest{
		TaskId:     taskID,
		TaskType:   taskType,
		Input:      input,
		Parameters: params,
	}

	log.Printf("[Master] 分配任务 %s 给 Agent %s", taskID, agent.AgentID)

	stream, err := agent.Client.ExecuteTask(ctx, req)
	if err != nil {
		return err
	}

	// 接收流式响应
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("[Master] 任务 %s [%s]: %d%% - %s",
			taskID, agent.AgentName, resp.Progress, resp.Content)
	}

	return nil
}

// 启动 Master gRPC 服务
func (s *MasterServer) startGRPC() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	grpcServer := grpc.NewServer()
	masterpb.RegisterMasterServiceServer(grpcServer, s)

	log.Println("[Master] 服务启动在 :50050")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

// 健康检查协程
func (s *MasterServer) healthCheckLoop() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.agentMutex.Lock()
		now := time.Now()
		for id, agent := range s.agents {
			agent.Mutex.Lock()
			if now.Sub(agent.LastHeartbeat) > 10*time.Second {
				agent.Status = "unhealthy"
				log.Printf("[Master] ⚠️ Agent %s 已失联", id)
			}
			agent.Mutex.Unlock()
		}
		s.agentMutex.Unlock()
	}
}

func (s *MasterServer) startTaskScheduler() {
	type Task struct {
		ID    string
		Type  string
		Input string
	}
	// 任务队列
	taskQueue := []Task{
		{ID: "task-001", Type: "计算", Input: "计算 1+1"},
		{ID: "task-002", Type: "分析", Input: "分析数据样本"},
		{ID: "task-003", Type: "处理", Input: "处理图片集"},
		{ID: "task-004", Type: "流式处理", Input: "流式处理日志"},
		{ID: "task-005", Type: "计算", Input: "复杂计算"},
	}

	pendingTasks := make(map[string]Task)
	for _, task := range taskQueue {
		pendingTasks[task.ID] = task
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for len(pendingTasks) > 0 {
		select {
		case <-ticker.C:
			// 尝试分发待处理任务
			for taskID, task := range pendingTasks {
				agent, err := s.SelectAgent(task.Type)
				if err != nil {
					log.Printf("任务 %s 暂无可用 Agent，等待重试", taskID)
					continue
				}

				log.Printf("分配任务 %s 给 Agent %s", taskID, agent.AgentName)
				go func(t Task, a *AgentInfo) {
					err := s.ExecuteTaskOnAgent(a, t.ID, t.Type, t.Input, nil)
					if err == nil {
						// 任务成功分发，从待处理列表中移除
						s.agentMutex.Lock()
						delete(pendingTasks, t.ID)
						s.agentMutex.Unlock()
						log.Printf("✅ 任务 %s 已分发", t.ID)
					} else {
						log.Printf("❌ 任务 %s 分发失败: %v", t.ID, err)
					}
				}(task, agent)
			}
		}
	}

	log.Println("✅ 所有任务已分发完成")
}

func main() {
	master := NewMasterServer()

	// 启动 Master 服务
	go master.startGRPC()

	// 启动健康检查
	go master.healthCheckLoop()

	// 等待 Agent 注册
	time.Sleep(2 * time.Second)

	// 启动任务调度器（会持续尝试）
	go master.startTaskScheduler()

	// 保持运行
	select {}
}
