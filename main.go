package main

import (
	"bwastartup/auth"
	"bwastartup/config"
	"bwastartup/controller"
	"bwastartup/user"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := config.Database()

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	token,err:=authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxOX0.2fE63FXAZaNE3t-gnis-rvU8mBQrSaO1HpPO752SILw")
	if err!=nil {
		fmt.Println("Lagi Error")
	}
	if token.Valid {
		fmt.Println("Token Valid")
	}
	userController := controller.NewUserHandler(userService,authService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/session", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatar", userController.UploadAvatar)
	router.Run()
}
