package model

import (
	"time"
)

type User struct {
	ID        uint      `Gorm:"primary" json:"id"`
	Username  string    `Gorm:"type:varchar(64);unique" json:"username"`
	Password  string    `Gorm:"type:varchar(255)" json:"password"`
	Age       int       `Gorm:"type:int" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
