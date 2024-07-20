package user

import "time"

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Occupation     string `json:"occupation"`
	PasswordHash   string `json:"password_hash"`
	AvatarFileName string `json:"avatar_file_name"`
	Role           string `json:"role"`
	CreatedAt      time.Time
	UpdatedAt	   time.Time
}