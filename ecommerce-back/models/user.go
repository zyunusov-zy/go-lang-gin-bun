package models

import "github.com/uptrace/bun"

// User model to represent users in the 'users' table
type User struct {
	bun.BaseModel `bun:"table:users"` // Correctly mark the table name for Bun

	ID       int64  `bun:"id,pk,autoincrement"` // Primary key with autoincrement
	Name     string `json:"name"`                // JSON tag for Name
	Email    string `json:"email,unique"`         // JSON tag for Email
	Password string `json:"password"`                    // Ignore the Password in JSON responses
	Role     string `json:"role"`                 // JSON tag for Role
}
