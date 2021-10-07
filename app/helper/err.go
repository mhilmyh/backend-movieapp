package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func SendResponseError(c *gin.Context, err error)  {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
}