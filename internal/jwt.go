package internal

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var signKey string = "employee_token"

type Jwt struct {
	SigningKey []byte
}
//载荷
type Claims struct {
	Name string
	jwt.StandardClaims
}

func getSignKey() string {
	return signKey
}

//创建jwt
func NewJwt() *Jwt {
	return &Jwt{[]byte(getSignKey())}
}

//创建token
func (j *Jwt)CreateToken(cl Claims) (string,error) {
	cl.IssuedAt = int64(time.Now().Unix()) //签发时间
	cl.ExpiresAt = int64(time.Now().Add(time.Hour*12).Unix()) //过期时间
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,cl)
	return token.SignedString(j.SigningKey)
}

//解析token
func (j *Jwt)ParseToken(tokenStr string)(*Claims,error){
	token,err := jwt.ParseWithClaims(tokenStr,&Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey,nil
	})
	if err != nil {
		if ve,ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed !=0 {
				return nil,errors.New("token无效")
			} else if ve.Errors&jwt.ValidationErrorExpired !=0 {
				return nil,errors.New("token已过期")
			} else if ve.Errors&jwt. ValidationErrorNotValidYet != 0 {
				return nil,errors.New("token未验证")
			} else {
				return nil,errors.New("token错误")
			}
		}
	}
	if claims,ok := token.Claims.(*Claims); ok && token.Valid{
		return claims,nil
	}
	return nil,errors.New("token错误")
}