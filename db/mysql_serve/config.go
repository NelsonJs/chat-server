package mysql_serve

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	dsn := "root:123456@tcp(localhost:3310)/demo?charset=utf8mb4"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
}
