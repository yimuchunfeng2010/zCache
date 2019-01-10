package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestDecrBy(context *gin.Context) {
	key := context.Param("key")
	value := context.Param("value")
	err := DecrBy(key, value)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Status": "Fail", "Data": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"Status": "Success", "Data": ""})
	}
}
