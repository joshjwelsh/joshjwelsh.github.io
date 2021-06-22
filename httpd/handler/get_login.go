package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Html5 Article Engine",
		})
	}
}
