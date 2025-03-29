package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"movie/internal/models"
	"net/http"

	services "movie/internal/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(auth *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: auth,
	}
}

func (this *AuthController) SingUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := this.AuthService.SingUp(user)
	if err != nil {
		log.Printf("Error signing up: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (this *AuthController) Login(c *gin.Context) {
	var credentials models.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		log.Printf("Error binding json: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, token, err := this.AuthService.Login(credentials)
	if err != nil {
		log.Printf("Error login: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (this *AuthController) PromoteToAdmin(c *gin.Context) {
	var userID = c.Param("userId")
	if err := this.AuthService.PromoteToAdmin(userID); err != nil {
		log.Printf("Error promoting to admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully promoted to admin",
	})
}
