package seed

import (
	"log"
	"rian-anggara/car-rental-api/models"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []string{"admin", "customer"}

	for _, roleName := range roles {
		var r models.Role
		result := db.Where("name =?", roleName).First(&r)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				// Role not found, create it
				err := db.Create(&models.Role{Name: roleName}).Error
				if err != nil {
					log.Printf("Failed to seed role '%s': %v", roleName, err)
				} else {
					log.Printf("Seeded role: %s", roleName)
				}
			} else {
				log.Printf("Error querying role '%s': %v", roleName, result.Error)
			}
		} else if result.RowsAffected == 0 {
			// No role found, create it
			err := db.Create(&models.Role{Name: roleName}).Error
			if err != nil {
				log.Printf("Failed to seed role '%s': %v", roleName, err)
			} else {
				log.Printf("Seeded role: %s", roleName)
			}
		} else {
			log.Printf("Role '%s' already exists, skipping", roleName)
		}
	}
}
