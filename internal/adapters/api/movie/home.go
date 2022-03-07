package api

import "github.com/gin-gonic/gin"

func (h *HTTPHandler) Home() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Welcome to my Movie-Api",
		})
	}
}
