package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"crowdfunding-server/handlers"
	"crowdfunding-server/models"
	"crowdfunding-server/repositories"
	"crowdfunding-server/services"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error!")
	}

	// Migration model
	db.AutoMigrate(&models.User{})

	userRepository := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService()

	userHandler := handlers.NewUserHandler(userService, authService)

	router := gin.Default()

	// Api version
	v1 := router.Group("/api/v1")

	v1.POST("/users", userHandler.Create)
	v1.POST("/users/login", userHandler.Login)
	v1.POST("/users/email_checkers", userHandler.CheckEmailAvailability)
	v1.POST("/users/upload_avatar", userHandler.UploadAvatar)

	router.Run()
}
