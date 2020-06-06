package models

import "github.com/jinzhu/gorm"

//User - struct
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
} 

//TableName - returns DB table name
func (user *User) TableName() string {
	return "user_credentials"
}
