package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestGetKeyNum(context *gin.Context) {
	num, err := GetKeyNum()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"status": "fail"})
	} else {
		context.JSON(http.StatusOK, gin.H{"value": num, "status": "done"})
	}
	return
}
