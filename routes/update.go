package routes

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Update(context *gin.Context){
	key := context.Param("key")
	value := context.Param("value")
	fmt.Println(key,value)
}

