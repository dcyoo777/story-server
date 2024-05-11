package router

import (
	"example/request"
	"example/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func UseSubRouter(r *gin.Engine, commonReq request.CommonRequestInterface, parseSelectQuery service.MakeQuery) {

	r.GET(
		fmt.Sprintf("/%s", commonReq.GetName()),
		func(c *gin.Context) {

			items, err := commonReq.Select(parseSelectQuery(c)...)

			fmt.Printf("ITEMS : %+v\n", items)
			fmt.Printf("ERROR : %+v\n", err)
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
			id, e1 := uuid.Parse(c.Params.ByName(commonReq.GetPrimaryKey()))
			if e1 != nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}
			item, err := commonReq.GetOne(id)
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

			id, e1 := uuid.Parse(c.Params.ByName(commonReq.GetPrimaryKey()))
			if e1 != nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}

			if err := c.BindJSON(&updatedItem); err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}

			fmt.Printf("%v %+v\n", id, updatedItem)

			item, err := commonReq.Update(id, updatedItem)

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
			id, e1 := uuid.Parse(c.Params.ByName(commonReq.GetPrimaryKey()))
			if e1 != nil {
				c.JSON(http.StatusNotFound, gin.H{})
				return
			}

			item, err := commonReq.Delete(id)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)
}
