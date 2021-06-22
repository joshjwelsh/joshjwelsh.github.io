package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	}
}
