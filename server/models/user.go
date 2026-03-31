package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email" gorm:"unique" gorm:"not null"`
	AvatarUrl   *string    `json:"avatarurl"`
	Password    string    `json:"password" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	AvatarUrl *string `json:"avatarurl,omitempty"`
}