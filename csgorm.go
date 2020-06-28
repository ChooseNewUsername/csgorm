package csgorm

import (
	"csgorm/log"
	"csgorm/session"
	"database/sql"
)

type Engine struct {
	db *sql.DB
}
func NewEngine(driver,source string)(e *Engine, err error)  {
	db,err := sql.Open(driver,source)

	if err!=nil {
		log.Error(err)
	}
	err = db.Ping();

	if err != nil {
		log.Error(err)
	}
	log.Info("sql link success!!")
	e = &Engine{db:db}
	return
}
func (e *Engine) Close()  {
	err :=  e.db.Close()
	if err != nil {
		log.Error(e)
	}
}

func (this *Engine)NewSession()(s *session.Session)  {
	return  session.New(this.db)
}