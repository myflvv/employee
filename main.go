package main

import (
	"employee/dao"
	"employee/internal"
	"employee/router"
	"fmt"
	"github.com/pkg/errors"
)
type ErrorTyped struct {
	error
}

func main()  {
	defer dao.DB.Close()
	internal.InitLog()
	r := router.SetupRouter()
	r.Run(":8881")
	//fmt.Printf("%+v",errors.New("test"))
	//err := error(ErrorTyped{errors.New("an error occurred")})
	//err = errors.Wrap(err, "wrapped")
	//fmt.Println("wrapped error: ", err)
	//fmt.Printf("%+v",errors.Cause(err))

	//switch errors.Cause(err).(type) {
	//case ErrorTyped:
	//	fmt.Println("a typed error occurred: ", err)
	//default:
	//	fmt.Println("an unknown error occurred")
	//}
	//StackTrace()
}


func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}
