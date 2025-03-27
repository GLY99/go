package routing

import (
	"ginWeb/controller"

	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.POST("/users", controller.CreateUser)
	v1.POST("/users/batch", controller.BatchCreateUsers)
	v1.DELETE("/users/:id", controller.DeleteUser)
	v1.GET("/users/:id", controller.GetUser)
	v1.GET("/users", controller.ListUsers)
	v1.PUT("/users/:id", controller.UpdateUser)
}
