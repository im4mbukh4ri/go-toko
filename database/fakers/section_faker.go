package fakers

import (
	"github.com/im4mbukh4ri/go-toko/app/models"
	"gorm.io/gorm"
)

func SectionFaker(db *gorm.DB) *models.Section {
	return &models.Section{}
}
