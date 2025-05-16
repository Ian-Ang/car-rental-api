package controllers

import (
	"net/http"

	"rian-anggara/car-rental-api/config"
	"rian-anggara/car-rental-api/dto"
	"rian-anggara/car-rental-api/models"

	"github.com/gin-gonic/gin"
)

// GET /users
//
//	@Summary		Get users
//	@Description	List of all user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserOutput
//	@Router			/users/ [get]
//	@security		BearerAuth
func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Role").Find(&users)

	var output []dto.UserOutput
	for _, user := range users {
		output = append(output, dto.UserOutput{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role.Name,
		})
	}

	c.JSON(http.StatusOK, output)
}

// GET /users/me
//
//	@Summary		Get My Profile
//	@Description	List detail my profile
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserOutput
//	@Failure		404
//	@Router			/users/me [get]
//	@security		BearerAuth
func GetMyProfile(c *gin.Context) {
	userID := c.GetString("userID")
	var user models.User

	if err := config.DB.Preload("Role").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	output := dto.UserOutput{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role.Name,
	}

	c.JSON(http.StatusOK, output)
}

// DELETE /users/:id (Admin Only)
//
//	@Summary		Delete users
//	@Description	Delete users (Admin Only) by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		200
//	@Failure		404
//	@Failure		500
//	@Router			/users/{id} [delete]
//	@security		BearerAuth
func DeleteUser(c *gin.Context) {
	user_id := c.Param("id")
	var user models.User

	// if err := config.DB.First(&user, user_id).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	// 	return
	// }
	if err := config.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GET DELETE USER /users/getdeleletuser(Admin Only)
//
//	@Summary		Get Delete users
//	@Description	Get all delete users (Admin Only)
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Failure		404
//	@Failure		500
//	@Router			/users/getdeleletuser [get]
//	@security		BearerAuth
func GetDeleteUser(c *gin.Context) {
	var users []models.User
	config.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&users)
	c.JSON(http.StatusOK, users)
}

// RESTORE USER /users/restore/:id (Admin Only)
//
//	@Summary		Restore users
//	@Description	Restore users (Admin Only) by user ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		200
//	@Failure		404
//	@Failure		500
//	@Router			/users/restore/{id} [post]
//	@security		BearerAuth
func RestoreUser(c *gin.Context) {
	user_ID := c.Param("id")
	config.DB.Unscoped().Model(&models.User{}).Where("id = ?", user_ID).Update("deleted_at", nil)
	c.JSON(http.StatusOK, gin.H{"message": "User restored successfully"})
}
