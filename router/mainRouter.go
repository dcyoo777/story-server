package router

import (
	"example/request"
	"example/service/appUser"
	"example/service/story"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var db = make(map[string]string)

func SetupRouter(dataSourceName string) *gin.Engine {
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

	storyRouter := story.CommonRequests{
		CommonRequests: request.CommonRequests{
			Name:           "story",
			PrimaryKey:     "story_id",
			DatasourceName: dataSourceName,
		},
	}

	UseSubRouter(r, storyRouter, story.MakeQuery)

	userRouter := appUser.CommonRequests{
		CommonRequests: request.CommonRequests{
			Name:           "appUser",
			PrimaryKey:     "user_id",
			DatasourceName: dataSourceName,
		},
	}

	UseSubRouter(r, userRouter, appUser.MakeQuery)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // appUser:foo password:bar
		"manu": "123", // appUser:manu password:123
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
