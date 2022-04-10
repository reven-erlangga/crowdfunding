package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"crowdfunding-server/handlers"
	"crowdfunding-server/helpers"
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
	db.AutoMigrate(&models.Campaign{})
	db.AutoMigrate(&models.CampaignImage{})
	db.AutoMigrate(&models.Transaction{})

	userRepository := repositories.NewUserRepository(db)
	campaignRepository := repositories.NewCampaignRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)

	userService := services.NewUserService(userRepository)
	campaignService := services.NewCampaignService(campaignRepository)
	transactionService := services.NewTransactionService(transactionRepository, campaignRepository)
	authService := services.NewAuthService()

	userHandler := handlers.NewUserHandler(userService, authService)
	campaignHandler := handlers.NewCampaignHandler(campaignService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	router := gin.Default()
	router.Static("/images", "./assets/images")

	// Api version
	v1 := router.Group("/api/v1")

	// User web service
	v1.POST("/users", userHandler.Create)                                                               // register
	v1.POST("/users/login", userHandler.Login)                                                          // login
	v1.POST("/users/email_checkers", userHandler.CheckEmailAvailability)                                // check email available
	v1.POST("/users/upload_avatar", authMiddleware(authService, userService), userHandler.UploadAvatar) // upload avatar

	// Campaign web service
	v1.GET("/campaigns", campaignHandler.GetCampaigns)                                                         // get campaigns
	v1.GET("/campaigns/:id", campaignHandler.GetCampaign)                                                      // get detail campaigns
	v1.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)            // post campaigns
	v1.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)         // update campaigns
	v1.POST("/campaigns/upload_images", authMiddleware(authService, userService), campaignHandler.UploadImage) // post campaigns

	// Transaction web service
	v1.GET("campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions) // get transactions

	router.Run()
}

func authMiddleware(authService services.AuthService, userService services.UserService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.ApiResponse("Unauthenticated", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrToken := strings.Split(authHeader, " ")

		if len(arrToken) == 2 {
			tokenString = arrToken[1]
		}

		token, err := authService.ValidationToken(tokenString)

		if err != nil {
			response := helpers.ApiResponse("Unauthenticated", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helpers.ApiResponse("Unauthenticated", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helpers.ApiResponse("Unauthenticated", http.StatusUnauthorized, "error", "")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		ctx.Set("currentUser", user)
	}
}
