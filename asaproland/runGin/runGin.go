package runGin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunGIn() {
	router := gin.Default()

	// GET请求
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// POST请求
	router.POST("/login", func(c *gin.Context) {
		// 获取POST请求中传递过来的username和password参数
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "user" && password == "password" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful!",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid username or password",
			})
		}
	})

	// 启动服务并监听8080端口
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
