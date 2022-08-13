package models

import (
	"fmt"
	"log"

	"github.com/Gopherlinzy/go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, //表名前缀，`Tag`的表名是`blog_Tag`
			SingularTable: true,        //使用单数表名，启用该选项，此时，`Tag`的表名是`blog_Tag`
		},
	})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	InitAuth()
	InitTag()
	InitArticle()
}
