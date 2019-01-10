package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestDeleteAll(context *gin.Context) {

	err := DeleteAll()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"status": "failure", "reason": err.Error()})
	} else {
		context.JSON(http.StatusOK, gin.H{"status": "success"})
	}

}
