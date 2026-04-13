package users

import "time"

type UserEntity struct {
	ID          int
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	CreatedAt   time.Time
}