package domain

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserId    int `gorm:"column:user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
