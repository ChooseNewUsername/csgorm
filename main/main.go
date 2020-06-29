package main

import (
	"csgorm"
)
import _ "github.com/go-sql-driver/mysql"
func main()  {
	engine,_:= csgorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	defer engine.Close()
	s := engine.NewSession()


	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
}
