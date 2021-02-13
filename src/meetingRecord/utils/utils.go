package utils

import (
	md52 "crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Txt struct {
	Project      string
	User         string
	Achievements string
	Goal         string
}

type Text struct {
	gorm.Model
	Txt
}

type Usr struct {
	Username string
	Password string
	Status   int
}

type User struct {
	gorm.Model
	Usr
}

var DB *gorm.DB
var err error

func InitMysql() {
	usr := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.Int("port")
	dbname := beego.AppConfig.String("dbname")

	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", usr, pwd, host, port, dbname, timeout)
	fmt.Println(dsn)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to link database, error = " + err.Error())
	} else {
		DB.AutoMigrate(&User{}, &Text{})
	}
}

func FormatTimeStamp(times time.Time) string {
	timestamp := times.Format("2006-01-02 15:04:05")
	return timestamp
}

func MD5(passwd string) string {
	data := []byte(passwd)
	md5 := md52.Sum(data)
	hexMd5 := fmt.Sprintf("%x", md5)
	return hexMd5
}
