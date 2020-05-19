package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

func init()  {
	//v := viper.New()
	viper.SetConfigFile(getCurrentPath()+"/config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error:%s", err))
	}
}

func GetString(key string) string  {
	chkKey(key)
	return viper.GetString(key)
}

func GetInt(key string) int  {
	chkKey(key)
	return viper.GetInt(key)
}

func GetBool(key string) bool  {
	chkKey(key)
	return viper.GetBool(key)
}

func chkKey(key string)  {
	if !viper.IsSet(key) {
		panic(fmt.Errorf("error:key %s not found",key))
	}
}

//获取当前文件路径
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}