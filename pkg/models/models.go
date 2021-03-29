package models

import (
	"errors"
	"time"
)

var ErrorNoRecord = errors.New("models: no match record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}
