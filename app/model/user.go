package model

import (
	"employee/dao"
	"employee/internal"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name string `form:"name" json:"name" binding:"required"`
	Pass string	`form:"pass" json:"pass" binding:"required"`
}
//登录
func (u *User)Login() (token string,err error) {
	var r User
	res := dao.DB.Select("pass").Where(&User{Name:u.Name}).First(&r)
	if res.Error != nil {
		return "",res.Error
	}
	isTrue := validatePassword(u.Pass,r.Pass)
	if isTrue == false {
		return "",errors.New("密码错误")
	} else {
		//生成jwt token
		t := internal.NewJwt()
		token := &internal.Claims{Name:u.Name}
		ResToken,err := t.CreateToken(*token)
		if err != nil {
			internal.Logger.Error(err.Error())
			return "",errors.New("token创建失败")
		}
		return ResToken,nil
	}
}

//生成密码
func (u *User)generatePassword() (encryptPass string,err error) {
	res,err := bcrypt.GenerateFromPassword([]byte(u.Pass),bcrypt.DefaultCost)
	if err != nil {
		internal.Logger.Error(err.Error())
		return "",err
	}
	encryptPass = string(res)
	return encryptPass,nil
}

//校验密码
func validatePassword(pass string, encryptPass string) (isTrue bool) {
	var err error
	if err = bcrypt.CompareHashAndPassword([]byte(encryptPass), []byte(pass)); err != nil {
		return false
	}
	return true
}