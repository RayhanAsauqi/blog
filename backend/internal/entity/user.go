package entity

import "time"

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username      string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"type:varchar(255);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Role string `gorm:"type:varchar(255)" json:"role"`
}