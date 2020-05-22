package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func ResErr(c *gin.Context,httpCode int,code int,msg string,err error)  {
	message := make(map[string]interface{})
	message["code"] = code
	message["msg"] = msg
	if err != nil {
		message["err"] = err.Error()
	}
	c.JSON(httpCode,message)
}

func ResJsonData(c *gin.Context,code int,msg string,data interface{})  {
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : msg,
		"data" : data,
	})
}