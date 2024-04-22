package router

import (
	"example/request"
	"example/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UseSubRouter(r *gin.Engine, commonReq request.CommonRequestInterface, parseSelectQuery service.MakeQuery) {

	r.GET(
		fmt.Sprintf("/%s", commonReq.GetName()),
		func(c *gin.Context) {

			items, err := commonReq.Select(parseSelectQuery(c)...)

			fmt.Printf("ITEMS : %+v\n", items)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": items})
			}
		},
	)

	r.GET(
		fmt.Sprintf("/%s/:%s", commonReq.GetName(), commonReq.GetPrimaryKey()),
		func(c *gin.Context) {
			pk := c.Params.ByName(commonReq.GetPrimaryKey())
			item, err := commonReq.GetOne(pk)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)

	r.POST(
		fmt.Sprintf("/%s", commonReq.GetName()),
		func(c *gin.Context) {
			var newItem interface{}
			if err := c.BindJSON(&newItem); err != nil {
				return
			}

			item, err := commonReq.Create(newItem)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)

	r.PUT(
		fmt.Sprintf("/%s/:%s", commonReq.GetName(), commonReq.GetPrimaryKey()),
		func(c *gin.Context) {
			var updatedItem interface{}

			pk := c.Params.ByName(commonReq.GetPrimaryKey())
			if err := c.BindJSON(&updatedItem); err != nil {
				return
			}

			fmt.Printf("%v %+v\n", pk, updatedItem)

			item, err := commonReq.Update(pk, updatedItem)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)

	r.DELETE(
		fmt.Sprintf("/%s/:%s", commonReq.GetName(), commonReq.GetPrimaryKey()),
		func(c *gin.Context) {
			pk := c.Params.ByName(commonReq.GetPrimaryKey())

			item, err := commonReq.Delete(pk)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)
}
