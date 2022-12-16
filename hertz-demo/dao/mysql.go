package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz_demo/config"
)

var Db *gorm.DB
var err error

func InitMysql() error {
	//dsn := fmt.Sprintf()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Conf.User, config.Conf.MySQLConfig.Password, config.Conf.MySQLConfig.Host,
		config.Conf.MySQLConfig.Port, config.Conf.MySQLConfig.DB)

	Db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Printf("数据库连接错误，请检查参数", err)
		return err
	}

	return nil
}
