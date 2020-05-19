package middleware

import (
	"employee/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuth(c *gin.Context)  {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		internal.ResErr(c,http.StatusUnauthorized,1,"token不能为空",nil)
		c.Abort()
		return
	}
	if s := strings.Split(token, " "); len(s) == 2 {
		token = s[1]
	}
	t := internal.NewJwt()
	res,err := t.ParseToken(token)
	if err != nil {
		internal.ResErr(c,http.StatusUnauthorized,1,err.Error(),nil)
		c.Abort()
		return
	}
	c.Set("token",res)
}