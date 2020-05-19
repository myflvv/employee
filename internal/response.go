package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func ResErr(c *gin.Context,httpCode int,code int,msg string,err error)  {
	c.JSON(httpCode,gin.H{
		"code" : code,
		"msg" :msg,
		"err" : err,
	})
}

func ResJsonData(c *gin.Context,code int,msg string,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}