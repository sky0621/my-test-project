// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package rdb

import (
	"database/sql"
)

type Content struct {
	ID   []byte
	Name string
}

type Course struct {
	ID    sql.NullString
	Name  string
	Level int32
}

type Program struct {
	ID        []byte
	Question  string
	Answer    string
	ContentID []byte
}
