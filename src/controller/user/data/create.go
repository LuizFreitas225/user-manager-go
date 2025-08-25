package data

import "time"

type InputUserOfCreate struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type OutputUserOfCreate struct {
	ID               int64     `json:"id"`
	CreateDate       time.Time `json:"create_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
}
