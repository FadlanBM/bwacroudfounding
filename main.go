package main

import (
	"bwastartup/config"
	"bwastartup/controller"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	db,err:=config.Database()

	if err!= nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

	userRepository:=user.NewRepository(db)
	userService:=user.NewService(userRepository)
	
	userController:=controller.NewUserHandler(userService)

	router:=gin.Default()
	api:=router.Group("/api/v1")


	api.POST("/users",userController.RegisterUser)
	router.Run()


	// var users []user.User

	// if err:=config.DB.Find(&users).Error;err!=nil{
	// 	log.Print("connect gagal")
	// }

	// for _, user := range users {
    //     log.Println(user.Name)
    // }

	// router:=gin.Default()
	// router.GET("/users",handler)
	// router.Run()
}
