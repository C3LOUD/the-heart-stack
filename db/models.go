// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
)

type Todo struct {
	ID         int64
	Content    string
	IsFinished sql.NullBool
	CreatedAt  sql.NullTime
}
