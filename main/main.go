package main

import (
	"csgorm"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"
func main()  {
	engine,_:= csgorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	defer engine.Close()
	s := engine.NewSession()
	r,_ := s.Exec("DROP TABLE IF EXISTS User;")
	fmt.Print(r.RowsAffected())

}
