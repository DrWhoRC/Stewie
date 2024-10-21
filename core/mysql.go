package core

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dbconn *gorm.DB

func InitMysql() (dbconn *gorm.DB) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("root:wpywatsendw0517@tcp(%s:3306)/fim_v1?charset=utf8&parseTime=True&loc=Local", mysqlHost)
	//dsn := "root:wpywatsendw0517@tcp(127.0.0.1:3306)/fim_v1?charset=utf8&parseTime=True&loc=Local"
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("faied to connect db")
	} else {
		fmt.Println("connected", dbconn)
	}
	Dbconn = dbconn

	return dbconn
}
