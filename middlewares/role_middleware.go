package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OnlyAdmin restricts access to users with role = "admin"
func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// OnlyCustomer restricts access to users with role = "customer"
func OnlyCustomer() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "customer" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Customer access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// AllowRoles restricts access to one or more allowed roles
func AllowRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		c.Abort()
	}
}
