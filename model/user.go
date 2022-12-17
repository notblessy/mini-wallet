package model

import (
	"time"
)

// UserRepository :nodoc:
type UserRepository interface {
	Create(user *User) error
}

// User :nodoc:
type User struct {
	CustomerXid string    `json:"customer_xid" gorm:"primary_key"`
	CreatedAt   time.Time `gorm:"<-:create" json:"created_at"`
}

// User :nodoc:
type UserResponse struct {
	Token string `json:"token"`
}

func (u *User) NewResponse(t string) *UserResponse {
	return &UserResponse{
		Token: t,
	}
}
