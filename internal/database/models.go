// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
	Name      string
	Url       string
	UserID    uuid.NullUUID
}

type GooseDbVersion struct {
	ID        int32
	VersionID int64
	IsApplied bool
	Tstamp    time.Time
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
	Name      string
	ApiKey    string
	Lastname  sql.NullString
}
