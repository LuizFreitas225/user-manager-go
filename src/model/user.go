package model

import "time"

const UserTableName = "users"

type User struct {
	ID               int64     `db:"id" json:"id"`
	CreateDate       time.Time `db:"create_date" json:"create_date"`
	LastModifiedDate time.Time `db:"last_modified_date" json:"last_modified_date"`
	Name             string    `db:"name" json:"name"`
	Email            string    `db:"email" json:"email"`
	Password         string    `db:"password" json:"password"`
}
