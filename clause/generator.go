package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{})(string,[]interface{})

var generators map[Type]generator

func init()  {
	generators := make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
}
func _insert(values ...interface{})(string,[]interface{}) {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName,fields),[]interface{}{}

}
func _values(values ...interface{})(string,[]interface{}) {

	val := "VALUES ("+strings.Join(values[0].([]string), ",")+");"
	return val,[]interface{}{}
}