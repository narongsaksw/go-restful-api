package userontroller

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"data" : "users", 
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"data" : "register", 
	})
}