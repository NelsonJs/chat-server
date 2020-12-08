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
	dsn := "root:6678510Jk.@tcp(localhost:3306)/hometown?charset=utf8mb4"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
}
