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

func Casbin() (*casbin.Enforcer,error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.name"),
		config.GetString("db.charset"),
	)
	a, err := gormadapter.NewAdapter("mysql", dsn,true)
	if err != nil {
		return nil,err
	}
	e, err := casbin.NewEnforcer("config/rbac_with_domains_model.conf", a)
	if err != nil {
		return nil,err
	}
	err = e.LoadPolicy()
	if err != nil {
		return nil,err
	}
	return e,nil
}


func (ca *CasbinStruct)AddCasbin() (bool,error) {
	e ,err := Casbin()
	if err != nil {
		return false,err
	}
	r,err := e.AddPolicy(ca.RoleName,ca.Domain,ca.Path,ca.Method)
	if err != nil {
		return false,err
	}
	return r,nil
}

func (ca *CasbinStruct)CheckCasbin() (bool,error) {
	e ,err := Casbin()
	if err != nil {
		return false,err
	}
	r,err := e.Enforce(ca.RoleName,ca.Domain,ca.Path,ca.Method)
	if err != nil {
		return false,err
	}
	return r,nil
}