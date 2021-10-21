package userontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sscarry2/ginapi/configs"
	"github.com/sscarry2/ginapi/models"
	"github.com/sscarry2/ginapi/utils"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	configs.DB.Find(&users)
	c.JSON(200, gin.H{
		"data" : users, 
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

	//check duplicate email
	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "duplicate email",})
		return
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

	var user models.User
	result := configs.DB.First(&user, id)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found.",})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchUserByFullname(c *gin.Context) {
	fullname := c.Query("fullname")

	var users []models.User
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes(utils.Paginate(c)).Find(&users)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found.",})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}