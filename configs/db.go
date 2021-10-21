package configs

import (
	"fmt"
	"os"

	"github.com/sscarry2/ginapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("connect DB failed!")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("connect DB successfully!")

	//migration
	// db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{}, &models.Blog{})


	DB = db
}