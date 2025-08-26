package data

import "time"

type InputUserOfFindById struct {
	ID int `json:"id"`
}

type OutputUserOfFindById struct {
	ID               int    `json:"id"`
	CreateDate       time.Time `json:"create_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
}
