package clause

import "strings"

type clause struct {
	sql strings.Builder
	sqlVars strings.Builder
}

type Type int
const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
)
type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}
func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.sqlVars[name] = vars
}