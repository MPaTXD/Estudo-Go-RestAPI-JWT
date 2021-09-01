package migrations

import (
	"example.com/estudo/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Produto{})
	db.AutoMigrate(models.User{})
}
