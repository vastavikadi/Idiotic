package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Bio         *string    `json:"bio"`
	AvatarUrl   *string    `json:"avatarurl"`
	IsVerified  string    `json:"isverified"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
