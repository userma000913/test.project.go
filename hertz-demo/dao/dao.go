package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hertz_demo/conf"
	"log"
)

var dao *Dao

type Dao struct {
	c     *conf.AppConfig
	mysql *Mysql
}
type Mysql struct {
	*gorm.DB
}

func New(c *conf.AppConfig) *Dao {
	// 单例模式
	// todo 项目启动全局加载，下一版本考虑如何按需加载
	if dao == nil {
		dao = &Dao{
			c: c,
			// 初始化mysql
			mysql: initMysql(c.MySQLConfig),
		}
	}
	return dao
}

func initMysql(c *conf.MySQLConfig) *Mysql {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User, c.Password, c.Host,
		c.Port, c.DB)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println()
		return nil
	}
	return &Mysql{
		db,
	}
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}
