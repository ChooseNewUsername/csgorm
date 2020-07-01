package session

import (
	"csgorm/log"
	"csgorm/schema"
	"database/sql"
	"fmt"
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
	s.sql.WriteString("CREATE TABLE ` ")
	s.sql.WriteString(sch.Name)
	s.sql.WriteString("` (")
	for i:=0;i<len(sch.FieldNames);i++{
		f := sch.FieldMap[sch.FieldNames[i]]
		tag := f.Tag
		kov := strings.Split(tag, ";")
		fmt.Println(kov)
	}

	return s
}
func analysisTableCloumn(tag string) string {
	var sql string
	kov := strings.Split(tag, ";")
	for i:=0;i<len(kov);i++{
		if strings.Contains(kov[i], ":"){
			kv := strings.Split(kov[i], ":")

		}
	}
	return sql
}
