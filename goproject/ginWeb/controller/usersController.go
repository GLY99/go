package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ginWeb/db"
	"ginWeb/model"
)

func CreateUser(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start create user")
	var user model.User
	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}

	// 将controller层的model转成dao层的model,然后将数据写入db中
	err := db.DbHandler.CreateUser(ctx, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   0,
		"error":    "",
		"username": user.Name,
	})
}

func BatchCreateUsers(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start batch create users")
	var users []*model.User = []*model.User{}
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}

	err := db.DbHandler.CreateUsers(ctx, users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"error":  "",
	})
}

func DeleteUser(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start delete user")
	// 从 URL 路径参数中获取 id
	id := c.Param("id") // 注意：id 是字符串类型

	// 构建sql条件
	conditions := make([]db.Condition, 0)
	conditions = append(conditions, db.Condition{Field: "id", Operator: "=", Value: id})

	err := db.DbHandler.DeleteUsers(ctx, conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "error": ""})
}

func GetUser(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start get user")
	id := c.Param("id")

	conditions := make([]db.Condition, 0)
	conditions = append(conditions, db.Condition{Field: "id", Operator: "=", Value: id})

	users, err := db.DbHandler.GetUsers(ctx, conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"error":  "",
		"users":  users,
	})
}

func ListUsers(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start list users")
	queryParams := c.Request.URL.Query()

	// 保留所有值 ?color=red&color=blue
	multiParams := make(map[string][]string)
	for key, values := range queryParams {
		multiParams[key] = values
	}
	conditions := make([]db.Condition, 0)
	for field, values := range queryParams {
		operator := "="
		if len(values) > 1 {
			operator = "IN"
		}
		var tmpValues interface{} = values
		if len(values) == 1 {
			tmpValues = values[0]
		}
		conditions = append(conditions, db.Condition{Field: field, Operator: operator, Value: tmpValues})
	}
	users, err := db.DbHandler.GetUsers(ctx, conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"error":  "",
		"users":  users,
	})
}

func UpdateUser(c *gin.Context) {
	log := c.MustGet("logger").(*zap.Logger)
	ctx := context.WithValue(c.Request.Context(), "logger", log)
	log.Info("start update user")
	id := c.Param("id")
	var updates map[string]interface{} = make(map[string]interface{})
	// 虽然map存储的是指针,但是ShouldBindJSON设计上只能接收指针变量,所以这里需要传递指针
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	conditions := make([]db.Condition, 0)
	conditions = append(conditions, db.Condition{Field: "id", Operator: "=", Value: id})
	err := db.DbHandler.UpdateUsers(ctx, updates, conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 1, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "error": ""})
}
