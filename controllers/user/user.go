package userontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sscarry2/ginapi/configs"
	"github.com/sscarry2/ginapi/models"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"data" : "users", 
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Fullname: input.Fullname,
		Email: input.Email,
		Password: input.Password,
	}

	result := configs.DB.Create(&user)

	//create error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error,})
		return
	}

	c.JSON(201, gin.H{
		"message" : "register successfully", 
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