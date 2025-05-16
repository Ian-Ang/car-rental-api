package controllers

import (
	"net/http"
	"rian-anggara/car-rental-api/config"
	"rian-anggara/car-rental-api/dto"
	"rian-anggara/car-rental-api/models"
	"rian-anggara/car-rental-api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register
//
//	@Summary		Register
//	@Description	Add a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.RegisterInput	true	"Payload body [RAW]"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/auth/register [post]
func Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)

	db := config.DB
	var existing models.User
	if err := db.Where("email = ?", input.Email).First(&existing).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	var customerRole models.Role
	db.Where("name = ?", "customer").First(&customerRole)

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		RoleID:   customerRole.ID,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

// LoginUser
//
//	@Summary		Login
//	@Description	Login to Generate JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body	dto.LoginInput	true	"User credentials"
//	@Success		200
//	@Failure		400
//	@Failure		401
//	@Failure		500
//	@Router			/auth/login [post]
func Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB
	var user models.User
	if err := db.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role.Name,
		},
	})
}
