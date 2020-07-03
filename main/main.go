package main

import (
	"csgorm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type car struct {
	Id int `csgorm:"PRIMARY;int(10) unsigned NOT NULL AUTO_INCREMENT"`
	Name string `csgorm:"varchar(10)  NOT NULL "`
	Age string `csgorm:"int(10)  NOT NULL "`
}

func do(values ...interface{}) string{
	fields := strings.Join(values[1].([]string), ",")
	return fields
}

type User struct {
	Name string
	Age int
}
func main()  {



	engine,_:= csgorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8")
	defer engine.Close()
	s := engine.NewSession()
	_,_ = s.CreateTable(car{}).Exec()
	u1 := &User{Name: "Tom", Age: 18}
	u2 := &User{Name: "Sam", Age: 25}
	s.Insert(u1,u2)
	//_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	//_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()

}
