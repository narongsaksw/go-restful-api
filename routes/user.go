package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/sscarry2/ginapi/controllers/user"
)

func InitUserRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", usercontroller.GetAllUsers)

	routerGroup.POST("/register", usercontroller.Register)

	routerGroup.POST("/login", usercontroller.Login)

	routerGroup.GET("/:id", usercontroller.GetUserById)

	routerGroup.GET("/search", usercontroller.SearchUserByFullname)
}