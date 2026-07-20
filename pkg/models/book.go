package models

import (
	"github.com/jinzhu/gorm"
	"github.com/thisispramod/go-bookstore/config"
	"github.com/thisispramod/go-bookstore/pkg/models"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	db.AutoMigrate(&Book{})
}
