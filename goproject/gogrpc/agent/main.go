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

type AgentServer struct {
	pb.UnimplementedAgentServiceServer
	agentID      string
	agentName    string
	address      string
	maxTasks     int32
	capabilities []string // agent能力类型
	runningTasks int32
	load         int32
	mutex        sync.RWMutex
	masterAddr   string
}

func NewAgentServer(agentID, agentName, address, masterAddr string) *AgentServer {
	return &AgentServer{
		agentID:      agentID,
		agentName:    agentName,
		address:      address,
		masterAddr:   masterAddr,
		maxTasks:     5,
		capabilities: []string{"计算", "分析", "处理", "流式处理"},
	}
}

// ExecuteTask 流式执行任务
func (s *AgentServer) ExecuteTask(req *pb.TaskRequest, stream pb.AgentService_ExecuteTaskServer) error {
	s.mutex.Lock()
	s.runningTasks++
	s.load = s.runningTasks * 20
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		s.runningTasks--
		s.load = s.runningTasks * 20
		s.mutex.Unlock()
	}()

	taskID := req.TaskId
	log.Printf("[Agent %s] 收到任务: ID=%s, Type=%s", s.agentID, taskID, req.TaskType)

	// 模拟任务处理
	stages := []struct {
		content  string
		progress int32
	}{
		{"开始执行任务...", 10},
		{"正在初始化...", 25},
		{"处理数据中...", 45},
		{"分析结果...", 65},
		{"生成报告...", 85},
		{"任务完成！", 100},
	}

	for _, stage := range stages {
		time.Sleep(300 * time.Millisecond)

		resp := &pb.TaskResponse{
			TaskId:     taskID,
			Content:    fmt.Sprintf("[%s] %s", s.agentName, stage.content),
			Progress:   stage.progress,
			IsComplete: stage.progress == 100,
			Status:     "running",
		}

		if stage.progress == 100 {
			resp.Status = "success"
			resp.Content = fmt.Sprintf("[%s] 任务 %s 执行完成！", s.agentName, taskID)
		}

		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}

// SimpleTask 简单任务
func (s *AgentServer) SimpleTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	s.mutex.Lock()
	s.runningTasks++
	s.load = s.runningTasks * 20
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		s.runningTasks--
		s.load = s.runningTasks * 20
		s.mutex.Unlock()
	}()

	time.Sleep(500 * time.Millisecond)

	return &pb.TaskResponse{
		TaskId:     req.TaskId,
		Content:    fmt.Sprintf("[%s] 任务完成: %s", s.agentName, req.Input),
		Progress:   100,
		IsComplete: true,
		Status:     "success",
	}, nil
}

// HealthCheck 健康检查
func (s *AgentServer) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	status := "healthy"
	if s.runningTasks >= s.maxTasks {
		status = "busy"
	}

	return &pb.HealthResponse{
		AgentId:      s.agentID,
		Status:       status,
		Load:         s.load,
		RunningTasks: s.runningTasks,
		MaxTasks:     s.maxTasks,
	}, nil
}

// registerToMaster 向 Master 注册
func (s *AgentServer) registerToMaster() error {
	conn, err := grpc.Dial(s.masterAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("连接 Master 失败: %v", err)
	}
	defer conn.Close()

	client := masterpb.NewMasterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &masterpb.RegisterRequest{
		AgentId:      s.agentID,
		AgentName:    s.agentName,
		Address:      s.address,
		MaxTasks:     s.maxTasks,
		Capabilities: s.capabilities,
	}

	resp, err := client.RegisterAgent(ctx, req)
	if err != nil {
		return err
	}

	log.Printf("[Agent %s] 注册成功: %s", s.agentID, resp.Message)
	return nil
}

// heartbeatLoop 心跳循环
func (s *AgentServer) heartbeatLoop() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		conn, err := grpc.Dial(s.masterAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("[Agent %s] 连接 Master 失败: %v", s.agentID, err)
			continue
		}

		client := masterpb.NewMasterServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

		s.mutex.RLock()
		req := &masterpb.HeartbeatRequest{
			AgentId:      s.agentID,
			RunningTasks: s.runningTasks,
			Load:         s.load,
		}
		s.mutex.RUnlock()

		resp, err := client.Heartbeat(ctx, req)
		cancel()
		conn.Close()

		if err != nil {
			log.Printf("[Agent %s] 心跳失败: %v", s.agentID, err)
			continue
		}

		if resp.Action == "stop" {
			log.Printf("[Agent %s] Master 要求停止", s.agentID)
			// 这里可以优雅关闭
		}
	}
}

func main() {
	// 配置参数
	agentID := "agent-003"
	agentName := "智能助手-C"
	agentAddr := ":50053"
	masterAddr := "localhost:50050"

	agent := NewAgentServer(agentID, agentName, agentAddr, masterAddr)

	// 启动 gRPC 服务
	lis, err := net.Listen("tcp", agentAddr)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAgentServiceServer(grpcServer, agent)

	// 注册到 Master
	if err := agent.registerToMaster(); err != nil {
		log.Printf("注册失败: %v", err)
	}

	// 启动心跳
	go agent.heartbeatLoop()

	log.Printf("[Agent %s] 服务启动在 %s", agentID, agentAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
