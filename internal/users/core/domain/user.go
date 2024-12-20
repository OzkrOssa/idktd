package domain

import (
	"time"
)

// UserRole is an enum for user's role
type UserRole string

// UserRole enum values
const (
	RoleAdmin  UserRole = "ROLE_ADMIN"
	RoleAgent  UserRole = "ROLE_AGENT"
	RoleReader UserRole = "ROLE_READER"
)

// User is an entity that represents a user
type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
