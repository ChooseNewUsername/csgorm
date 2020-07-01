package main

import (
	"csgorm"
	_ "github.com/go-sql-driver/mysql"
)

type car struct {
	Id int `csgorm:"json:Username;xml:user_name"`
	Name string `csgorm:"json:Username;xml:user_name"`
	Age string `csgorm:"json:Username;xml:user_name"`
}

func main()  {


	engine,_:= csgorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	defer engine.Close()
	s := engine.NewSession()
	s.CreateTable(car{})

	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	//_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
}
