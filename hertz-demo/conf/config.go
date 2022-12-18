package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
)

type AppConfig struct {
	Env          string `yaml:"env" json:"env"`
	Name         string `json:"name"`
	Port         int    `json:"port"`
	*MySQLConfig `yaml:"mysql" json:"mysql"`
	*RedisConfig `yaml:"redis" json:"redis"`
}

// 定义mysql配置文件的结构体
type MySQLConfig struct {
	Host         string `mapstructure:"host" yaml:"host"`
	User         string `mapstructure:"user" yaml:"user"`
	Password     string `mapstructure:"password" yaml:"password"`
	DB           string `mapstructure:"dbname" yaml:"dbname"`
	Port         int    `mapstructure:"port" yaml:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns" yaml:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
}

// 定义redis配置文件的结构体
type RedisConfig struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Password string `mapstructure:"password" yaml:"password"`
	Port     int    `mapstructure:"port" yaml:"port"`
	DB       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"pool_size" yaml:"pool_size"`
}

type EsConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
}

func InitConfig() *AppConfig {
	n := NewNacosClient()
	n.InitNacosClient()
	config, err := n.GetNacosConfigContent()
	if err != nil {
		log.Println(err)
		return nil
	}

	var c *AppConfig
	if err := yaml.Unmarshal([]byte(config), &c); err != nil {
		log.Printf("Unmarshal is err;err=%s", err)
		return nil
	}
	fmt.Println("conf ok")
	return c
}
