package controller

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service,authService auth.Service) *userHandler {
	return &userHandler{userService: userService,authService: authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationInputError(err)
		errorMessage := gin.H{"errors": errors}
		res := helper.NewResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		res := helper.NewResponse("Register account failed", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err!=nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helper.NewResponse("Login account failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := user.FormatUsers(newUser, token)
	res := helper.NewResponse("Accunt has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		errors := helper.FormatValidationInputError(err)
		errorMessage := gin.H{"errors": errors}
		res := helper.NewResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	accounVerify, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helper.NewResponse("Login account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	token, err := h.authService.GenerateToken(accounVerify.ID)
	if err!=nil {
		errorMessage := gin.H{"errors": err.Error()}
		res := helper.NewResponse("Login account failed", http.StatusInternalServerError, "error", errorMessage)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	formatter := user.FormatUsers(accounVerify, token)
	res := helper.NewResponse("Accunt has been Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil {
		errors := helper.FormatValidationInputError(err)
		errorMessage := gin.H{"errors": errors}
		res := helper.NewResponse("Email Checking account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		res := helper.NewResponse("Email Checking account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	fmt.Print(isEmailAvailable)

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email tidak tersedia"
	if isEmailAvailable {
		metaMessage = "Email sudah tersedia"
	}

	res := helper.NewResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	imagePath := viper.GetString("IMAGE_STORAGE_PATH")

	file, err := c.FormFile("avatar")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		res := helper.NewResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	userId := 2

	userIdStr := strconv.Itoa(userId)
	currentTime := time.Now().Format("20060102150405")
	nameImage := fmt.Sprintf("%s-%s-%s", userIdStr, currentTime, file.Filename)
	imageSave := imagePath + nameImage
	err = c.SaveUploadedFile(file, imageSave)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		res := helper.NewResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, err = h.userService.SaveAvatar(userId, nameImage)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}
		res := helper.NewResponse("Failed to upload avatar image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}

	res := helper.NewResponse("Avatar success update", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, res)
}
