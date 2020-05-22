package router

import (
	"employee/app/middleware"
	"employee/app/servcie"
	"employee/dao"
	"employee/schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/login", servcie.Login)
		v1.GET("/tt", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"code" : 1,
				"msg":"eee",
			})
		})
		v1.POST("/permission/add", servcie.AddPermission)
		v1.GET("/mr", func(c *gin.Context) {
			dao.DB.AutoMigrate(&schema.User{})
		})
		v1.Use(middleware.JwtAuth)
		v1.GET("/test", func(c *gin.Context) {
			token := c.Request.Header.Get("Authorization")
			fmt.Println(token)
		})
	}
	return r
}
