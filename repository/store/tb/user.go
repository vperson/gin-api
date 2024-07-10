package tb

import "gorm.io/gorm"

// User 用户表
type User struct {
	Model
	Username     string `json:"username" gorm:"uniqueIndex;size:100"`
	Phone        string `json:"phone" gorm:"uniqueIndex;size:20"`
	Password     string `json:"password"`
	IP           string `json:"ip"`
	LoginFailure uint   `json:"login_failure"`
}

type UserDB struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}
