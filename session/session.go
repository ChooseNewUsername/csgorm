package session

import (
	"csgorm/clause"
	"database/sql"
	"strings"
)

type Session struct {
	db *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
	clause   clause.Clause
}

func New(db *sql.DB)(*Session)  {
	return &Session{db:db}
}

func (s *Session)Clear()  {
	s.sql.Reset();
	s.sqlVars = nil
	s.clause = clause.Clause{}
}