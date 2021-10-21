package userontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/matthewhartstonge/argon2"
	"github.com/sscarry2/ginapi/configs"
	"github.com/sscarry2/ginapi/models"
	"github.com/sscarry2/ginapi/utils"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	configs.DB.Preload("Blogs").Find(&users)
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
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Email: input.Email,
		Password: input.Password,
	}

	//check email
	userAccount := configs.DB.Where("email = ?", input.Email).First(&user)
	if userAccount.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found.",})
		return
	}

	//verify password
	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized."})
		return
	}

	//create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	accessToken, _ := claims.SignedString([]byte(jwtSecretKey))

	c.JSON(201, gin.H{
		"message" : "login successfully",
		"accessToken": accessToken, 
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
	c.JSON(http.StatusOK, gin.H{
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

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}