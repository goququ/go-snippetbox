package models

import (
	"errors"
	"time"
)

var (
	ErrorNoRecord         = errors.New("models: no match record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
	Created  time.Time
	Active   bool
}
