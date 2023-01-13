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

type Content struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Thumbnail   string
	Content     string
	Description string
	Playtime    time.Time
	Resolution  string
	Aspect      string
	Tag         bool
	Category    string
	UserId      int `gorm:"column:user_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
