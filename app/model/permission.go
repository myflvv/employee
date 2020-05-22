package model

import "employee/internal"

type Permission struct {
	RoleName string `form:"rolename" binding:"required"`
	Path string	`form:"path" binding:"required"`
	Domain string `form:"domain" binding:"required"`
	Method string `form:"method" binding:"required"`
}

func (p *Permission)AddPermission() (bool,error)  {
	ca := internal.CasbinStruct{RoleName:p.RoleName,Domain:p.Domain,Path:p.Path,Method:p.Method}
	//a.RoleName = "admin"
	//a.Domain = "s.com"
	//a.Path = "/test/2"
	//a.Method = "GET"
	//a.RoleName = ca.RoleName
	r,err := ca.AddCasbin()
	if err != nil {
		internal.Logger.Error(err.Error())
		return false,err
	}
	return r,nil
}

func Checkp()  {
	a := internal.CasbinStruct{}
	a.RoleName = "admin"
	a.Domain = "s.com"
	a.Path = "/test/2"
	a.Method = "GET"

}