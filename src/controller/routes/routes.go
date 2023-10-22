package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lechitz/GoMongoConnect/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserId/:userId", controller.FindUserByID)
	r.GET("/getUserEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
