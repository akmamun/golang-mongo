package models

import (
	"time"
)

//Todo model
type Todo struct {
	Title      *string   `json:"title" validate:"required,min=2,max=100"`
	Body       *string   `json:"body" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
