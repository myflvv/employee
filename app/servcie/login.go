package servcie

import (
	"employee/app/model"
	"employee/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	*model.User
}

type token struct {
	Token string `json:"token"`
}

func Login(c *gin.Context)  {
	var p User
	if err := c.ShouldBind(&p);err != nil {
		internal.ResErr(c,http.StatusBadRequest,1,"用户名或密码不能为空",nil)
		return
	}
	tokens,err := p.Login()
	if err != nil {
		internal.ResErr(c,http.StatusBadRequest,1,err.Error(),nil)
		return
	} else {
		internal.ResJsonData(c,0,"success", &token{Token:tokens})
		return
	}
}