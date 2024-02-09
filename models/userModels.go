package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey;"`
	Name        string
	Email       string `gorm:"unique;"`
	Password    string
	OTP         string
	// Ball        bool
	Language    string `gorm:"default:en_US"`
	IsSuperuser bool   `gorm:"default:false"`
	Is2FA       bool   `gorm:"default:true"`
}