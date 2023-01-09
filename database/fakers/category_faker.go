package fakers

import (
	"log"

	"github.com/im4mbukh4ri/go-toko/app/models"
	"gorm.io/gorm"
)

func Categoryaker(db *gorm.DB) *models.Category {
	section := SectionFaker(db)
	err := db.Create(&section).Error
	if err != nil {
		log.Fatal(err)
	}
	return &models.Category{
		// ID        :uuid.New().String(),
		// SectionID :section.ID,
		// Products  :[]Product `gorm:"many2many:product_categories;"`
		// Name      string    `gorm:"size:100;"`
		// Slug      string    `gorm:"size:100;"`
		// CreatedAt time.Time
		// UpdatedAt time.Time
	}
}
