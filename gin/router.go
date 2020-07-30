package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// QueryString
	// curl -X GET "http://localhost:8080/welcome?firstname=Lee&lastname=Cheol"
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// PostForm
	// curl -X POST http://localhost:8080/form_post -d message="exist" -d nick="ben"
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "Guest")

		c.JSON(200, gin.H{
			"message": message,
			"nick":    nick,
		})
	})

	// Query + PostForm
	// curl -X POST "http://localhost:8080/post?id=wks0968&page=3" -d name="cheolhee" -d message="golang is fantastic"
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.Query("page")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.String(200, "id : %s\npage : %s\nname : %s\nmessage : %s", id, page, name, message)
	})

	router.Run(":8080")
}
