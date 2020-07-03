package session

import (
	"csgorm/clause"
	"go/ast"
	"reflect"
	"strings"
)

func(s *Session) Insert(model interface{}) *Session {
	s.clause.Set(clause.INSERT,)
	return s
}
func InsertRecord(model interface{}) string  {
	var sql strings.Builder
	var  field []string
	var  values []interface{}
	modelType := reflect.TypeOf(model)
	tableName := modelType.Name()
	sql.WriteString("INSERT INTO "+tableName)
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		values = append(values,modelType.FieldByName(p.Name))
		if !p.Anonymous && ast.IsExported(p.Name) {
			field = append(field, strings.ToLower(p.Name))
		}

	}
	fieldSql := strings.Join(field,",")

	sql.WriteString(" ("+fieldSql+") VALUES ("+)

	return sql.String()
}