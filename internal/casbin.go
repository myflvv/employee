package internal

import (
	"employee/config"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
)

type CasbinStruct struct {
	ID uint
	Ptype string
	RoleName string
	Path string
	Domain string
	Method string
}

func Casbin() *casbin.Enforcer {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.name"),
		config.GetString("db.charset"),
	)
	a, _ := gormadapter.NewAdapter("mysql", dsn,true)
	e, _ := casbin.NewEnforcer("config/rbac_with_domains_model.conf", a)
	e.LoadPolicy()
	return e
}

func (ca *CasbinStruct)Add(cs CasbinStruct) (bool) {
	e := Casbin()
	r,err := e.AddPolicy(cs.RoleName,cs.Path,cs.Method,cs.Domain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}