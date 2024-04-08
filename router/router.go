package router

import (
	"example/service"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "custom-header", "content-type"},
			AllowCredentials: true,
			MaxAge:           1 * time.Minute,
		}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// GetAll Story
	r.GET("/story", func(c *gin.Context) {
		items, err := service.StoryCommonReq.GetAll()
		fmt.Printf("%v\n", items)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": items})
		}
	})

	// GetOne Story
	r.GET("/story/:story_id", func(c *gin.Context) {
		storyId := c.Params.ByName("story_id")
		item, err := service.StoryCommonReq.GetOne(storyId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": item})
		}
	})

	r.POST("/story", func(c *gin.Context) {
		var newStory service.Story
		if err := c.BindJSON(&newStory); err != nil {
			return
		}

		item, err := service.StoryCommonReq.Create(newStory.ToDB())

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": item})
		}
	})

	r.PUT("/story/:story_id", func(c *gin.Context) {
		var updatedStory service.Story

		storyId := c.Params.ByName("story_id")
		if err := c.BindJSON(&updatedStory); err != nil {
			return
		}

		item, err := service.StoryCommonReq.Update(storyId, updatedStory)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": item})
		}
	})

	r.DELETE("/story/:story_id", func(c *gin.Context) {
		storyId := c.Params.ByName("story_id")

		item, err := service.StoryCommonReq.Delete(storyId)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": item})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* main curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
