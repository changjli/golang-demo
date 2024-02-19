package routes

import (
	atom_user "user_service/atom/user/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	// bisa langsung atau disini
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	user := route.Group("user")
	{
		user.GET("get/all", atom_user.GetAllData)
		user.POST("add", atom_user.CreateUser)
		user.POST("get/one", atom_user.GetUserByUsername)
		user.GET("get/:username", atom_user.GetUserByUsernameV2)
		user.PUT("update", atom_user.UpdateUserById)
		user.PUT("update/:id", atom_user.UpdateUserByIdV2)
		user.DELETE(":id", atom_user.DeleteUserById)
		user.DELETE("/delete", atom_user.DeactivateUserById)
	}

	return route
}
