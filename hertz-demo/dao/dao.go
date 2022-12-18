package dao

import (
	"hertz_demo/conf"
	"hertz_demo/proxy"
)

var dao *Dao

type Dao struct {
	c     *conf.AppConfig
	mysql *proxy.Mysql
}

func New(c *conf.AppConfig) *Dao {
	// 单例模式
	// todo 项目启动全局加载，下一版本考虑如何按需加载
	if dao == nil {
		dao = &Dao{
			c: c,
			// 初始化mysql
			//mysql: proxy.InitMysql(c.MySQLConfig),
		}
	}
	return dao
}

// Close release resource
func (d *Dao) Close() error {
	return nil
}
