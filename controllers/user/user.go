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

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"data": id,
	})
}

func SearchUserByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}