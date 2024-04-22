package story

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SelectQuery struct {
	StartFrom string `form:"startFrom"`
	StartTo   string `form:"startTo"`
}

func (query SelectQuery) ToExp() []exp.Expression {
	var result []exp.Expression
	if len(query.StartFrom) > 0 && len(query.StartTo) > 0 {
		result = append(result, goqu.C("start").Between(goqu.Range(query.StartFrom, query.StartTo)))
	}
	return result
}

func MakeQuery(c *gin.Context) []exp.Expression {
	var queries SelectQuery

	if c.ShouldBind(&queries) != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return []exp.Expression{}
	}

	return queries.ToExp()
}
