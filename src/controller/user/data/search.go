package data

import "time"

type InputUserOfSearch struct {
	SearchTerm string `json:"search_term"`
}

type OutputUserOfSearch struct {
	ID               int64     `json:"id"`
	CreateDate       time.Time `json:"create_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
}
