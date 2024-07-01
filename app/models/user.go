package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `json:"username" valid:"required~Username is required"`
	Email     string    `json:"email" gorm:"unique" valid:"required~Email is required,email~Invalid email format"`
	Password  string    `json:"-" valid:"required~Password is required, length(6|30)~Password must be between 6 to 30 characters"`
	Photo     *Photo    `json:"photo,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserPass struct {
	OldPassword string `json:"old_password" valid:"required~Old password is required"`
	NewPassword string `json:"new_password" valid:"required~New password is required"`
}
