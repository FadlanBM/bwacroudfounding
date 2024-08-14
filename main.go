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
	userController := controller.NewUserHandler(userService,authService)

	fmt.Println(authService.GenerateToken(34))
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userController.RegisterUser)
	api.POST("/session", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.POST("/avatar", userController.UploadAvatar)
	router.Run()
}
