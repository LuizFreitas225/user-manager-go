package routes

import (
	"github.com/LuizFreitas225/user-manager-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitiRoutes(router *gin.RouterGroup) {
	router.GET("/getUserById/:userId", controller.FindUserById)
	router.POST("/getUserByEmail", controller.FindUserByEmail)
	router.POST("/createUser", controller.CreateUser)
	router.PUT("/updateUser/:userId", controller.UpdateUser)
	router.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
