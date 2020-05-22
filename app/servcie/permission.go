package servcie

import (
	"employee/app/model"
	"employee/internal"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Permission struct {
	*model.Permission
}

func AddPermission(c *gin.Context)  {
	var p Permission
	if err := c.ShouldBind(&p); err != nil {
		internal.ResErr(c,http.StatusBadRequest,1,"参数不全",err)
		return
	}
	fmt.Println("test")

}