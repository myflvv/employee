package dao

import (
	"employee/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.GetString("db.user"),
		config.GetString("db.pass"),
		config.GetString("db.host"),
		config.GetString("db.port"),
		config.GetString("db.name"),
		config.GetString("db.charset"),
	)
	DB, err = gorm.Open("mysql", dsn)
	DB.SingularTable(true) //禁用表名复数
	DB.LogMode(true)
	gorm.DefaultTableNameHandler = func (DB *gorm.DB, defaultTableName string) string  {
		return config.GetString("db.table_prefix") + defaultTableName
	}
	//defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
}