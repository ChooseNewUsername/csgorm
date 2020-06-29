package session

import (
	"csgorm/log"
	"database/sql"
)

func(s *Session) Exec()(res sql.Result,err error)   {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if res, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string,values ...interface{}) *Session {
	s.sql.WriteString(sql);
	s.sql.WriteString(" ");
	s.sqlVars = append(s.sqlVars ,values...)
	return s
}