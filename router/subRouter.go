package router

import (
	"example/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubRouter struct {
	Path       string
	PrimaryKey string
}

func (subRouter SubRouter) useCommonRequests(r *gin.Engine, commonReq request.CommonRequestInterface) {
	// GetAll Story
	r.GET(
		fmt.Sprintf("/%s", subRouter.Path),
		func(c *gin.Context) {
			items, err := commonReq.GetAll()
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": items})
			}
		},
	)

	// GetOne Story
	r.GET(
		fmt.Sprintf("/%s/:%s", subRouter.Path, subRouter.PrimaryKey),
		func(c *gin.Context) {
			pk := c.Params.ByName(subRouter.PrimaryKey)
			item, err := commonReq.GetOne(pk)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)

	r.POST(
		fmt.Sprintf("/%s", subRouter.Path),
		func(c *gin.Context) {
			//var newStory service.Story
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
		fmt.Sprintf("/%s/:%s", subRouter.Path, subRouter.PrimaryKey),
		func(c *gin.Context) {
			//var updatedStory service.Story
			var updatedItem interface{}

			pk := c.Params.ByName(subRouter.PrimaryKey)
			if err := c.BindJSON(&updatedItem); err != nil {
				return
			}

			item, err := commonReq.Update(pk, updatedItem)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)

	r.DELETE(
		fmt.Sprintf("/%s/:%s", subRouter.Path, subRouter.PrimaryKey),
		func(c *gin.Context) {
			pk := c.Params.ByName(subRouter.PrimaryKey)

			item, err := commonReq.Delete(pk)

			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{"result": item})
			}
		},
	)
}
