package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"ID"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Picture  string `json:"picture"`
	Gender   string `json:"gender"`
}
