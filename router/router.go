package router

import (
	"example/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// GetAll Story
	r.GET("/story", func(c *gin.Context) {
		items, err := service.StoryCommonReq.GetAll()
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

//func main() {
//	r := setupRouter()
//	// Listen and Server in 0.0.0.0:8080
//	r.Run(":8080")
//}
