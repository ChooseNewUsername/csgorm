package session

import (
	"database/sql"
	"strings"
)

type Session struct {
	db *sql.DB
	sql strings.Builder
}

func New(db *sql.DB)(*Session)  {
	return &Session{db:db}
}
