package session

import (
	"database/sql"
	"strings"
)

type Session struct {
	db *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB)(*Session)  {
	return &Session{db:db}
}

func (s *Session)Clear()  {
	s.sql.Reset();
	s.sqlVars = nil
}