package session

import (
	"csgorm/log"
	"database/sql"
)

func(s *Session) Exec(sql string)(res sql.Result,err error)   {
	res,err = s.db.Exec(sql)
	if err != nil {
		log.Error(err)
	}
	return res,err
}
func Raw(sql string,vars ...interface{})  {

}