package data_source

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"shop/models/user"
)

var (
	Db *gorm.DB
	err error
)


func init()  {
	conf := LoadMysqlConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.UserName,
		conf.Pwd,
		conf.Host,
		conf.Port,
		conf.BaseName,
	)
	fmt.Println(dsn)
	Db,err = gorm.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	Db.AutoMigrate(user.User{})
}
