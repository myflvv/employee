package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)
func TestLogin(t *testing.T)  {
	Convey("登录",t, func() {
		u := &User{Name: "admin", Pass:"admin"}
		_,err := u.Login()
		So(err,ShouldBeNil)
	})
}
