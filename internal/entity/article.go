package entity

import "time"

type Article struct {
	ID        int64  `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(255); not null"`
	Content   string `gorm:"type:text; not null"`
	Author    string `gorm:"type:varchar(255); not null"`
	CreatedAt time.Time	
	UpdatedAt time.Time	
}