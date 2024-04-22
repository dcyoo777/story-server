package service

import (
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/gin-gonic/gin"
)

type MakeQuery func(c *gin.Context) []exp.Expression
