package fakers

import (
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/im4mbukh4ri/go-toko/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	// hash,err :=models.HashPassword("123456");
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      models.HashPassword("123456"),
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}

}
