package session

import (
	"csgorm/log"
	"csgorm/schema"
	"database/sql"
	"strings"
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

func (s *Session)CreateTable(dest interface{}) *Session  {
	sch := schema.Parse(dest)
	PRIMARY:=""
	s.sql.WriteString("CREATE TABLE ` ")
	s.sql.WriteString(sch.Name)
	s.sql.WriteString("` (")
	for i:=0;i<len(sch.FieldNames);i++{
		f := sch.FieldMap[sch.FieldNames[i]]
		tag := f.Tag
		kov := strings.Split(tag, ";")
		s.sql.WriteString("`"+strings.ToLower(f.Name)+"`  ")
		for i:=0;i<len(kov);i++{
			if kov[i] =="PRIMARY" {
				PRIMARY ="PRIMARY KEY (`"+strings.ToLower(f.Name)+"`)"
			}else {
				s.sql.WriteString(kov[i])
				s.sql.WriteString(" ")
			}

		}
		s.sql.WriteString(",")

	}
	s.sql.WriteString(PRIMARY)
	s.sql.WriteString(" )AUTO_INCREMENT=10207705 DEFAULT CHARSET=utf8;")
	return s
}

