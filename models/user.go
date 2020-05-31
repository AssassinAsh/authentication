package models

import "github.com/jinzhu/gorm"

//User - struct
type User struct {
	gorm.Model
	Username string
	Password string
	Phone    string
}

//TableName - returns DB table name
func (user *User) TableName() string {
	return "user_credentials"
}
