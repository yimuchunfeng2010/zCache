package routes

import (
	"github.com/gin-gonic/gin"
	"fmt")


func Get(context *gin.Context){
	key := context.Param("key")
	fmt.Println(key)
}
