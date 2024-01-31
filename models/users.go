package models

import "time"

type Users struct {
	ID       uint      `gorm:"primaryKey"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
