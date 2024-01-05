package models

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"` // Set field as primary key
	CreatedAt time.Time 
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}