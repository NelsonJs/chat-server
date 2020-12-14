package mysql_serve

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	port := viper.GetString("mysqlPort")
	pwd := viper.GetString("mysqlPwd")
	dbName := viper.GetString("mysqlDb")
	dsn := "root:"+pwd+"@tcp(localhost:"+port+")/"+dbName+"?charset=utf8mb4"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
}
