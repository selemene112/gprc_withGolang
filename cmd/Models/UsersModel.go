package Models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Email       string `gorm:"column:email"`
	Name        string `gorm:"column:name"`
	PhotoProfil string `gorm:"column:photo_profil"` // ubah ke snake_case
}
