package controllers

import (
	"net/http"

	"rian-anggara/car-rental-api/config"
	"rian-anggara/car-rental-api/dto"
	"rian-anggara/car-rental-api/models"

	"github.com/gin-gonic/gin"
)

// GET /roles
//
//	@Summary		Get role
//	@Description	list of all role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.RoleOutput
//	@Failure		400
//	@Failure		500
//	@Router			/roles/ [get]
//	@security		BearerAuth
func GetRoles(c *gin.Context) {
	var roles []models.Role
	config.DB.Find(&roles)

	var output []dto.RoleOutput
	for _, r := range roles {
		output = append(output, dto.RoleOutput{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, output)
}

// POST /roles
//
//	@Summary		Create role
//	@Description	Add a new role
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RoleInput	true	"Payload body [RAW]"
//	@Success		200		{object}	dto.RoleOutput
//	@Failure		400
//	@Failure		500
//	@Router			/roles/ [post]
//	@security		BearerAuth
func CreateRole(c *gin.Context) {
	var input dto.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := models.Role{Name: input.Name}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
		return
	}

	output := dto.RoleOutput{
		ID:   role.ID,
		Name: role.Name,
	}

	c.JSON(http.StatusCreated, output)
}

// PUT /roles/:id
//
//	@Summary		Update role
//	@Description	Update role by ID
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string		true	"Role ID"
//	@Param			request	body		dto.RoleInput	true	"Payload body [RAW]"
//	@Success		200
//	@Failure		400
//	@Failure		500
//	@Router			/roles/{id} [put]
//	@security		BearerAuth
func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	var input dto.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.Name = input.Name
	config.DB.Save(&role)

	c.JSON(http.StatusOK, gin.H{
		"id":   role.ID,
		"name": role.Name,
	})
}
