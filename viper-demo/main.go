package main

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/google/martian/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"time"
)

// viper
func main() {

}

type AppConfig struct {
}

var Conf = new(AppConfig)

// 读字节，可以是自定义的，也可以是网络请求来的
func writeByte() {
	viper.SetConfigType("json") // 支持 json、yaml、toml等
	byteList := []byte("自定义内容")
	viper.ReadConfig(bytes.NewBuffer(byteList))

	// 无法监听
}

// 读文件
func writeFile(path string, name string) error {
	// 读文件的时候必须给一个文件的名称或者地址
	viper.SetConfigFile(path)
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return nil
}

// 远程配置调用（仅etcd、consul）
func writeRemote() {

	// 添加依赖
	// _ "github.com/spf13/viper/remote"
	viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/conf/hugo.json")
	viper.SetConfigType("json")
	_ = viper.ReadRemoteConfig()

}

// 监听远程配置
func WatchRemote() {
	// alternatively, you can create a new viper instance.
	var runtime_viper = viper.New()

	runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/conf/hugo.yml")
	runtime_viper.SetConfigType("yaml") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

	// read from remote conf the first time.
	_ = runtime_viper.ReadRemoteConfig()

	// unmarshal conf
	runtime_viper.Unmarshal(&Conf)

	// open a goroutine to watch remote changes forever
	// 必须开一个协程后台监听，且必须是死循环监听，因为WatchRemoteConfig并没有死循环 而文件监听的WatchConfig有！！！
	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				log.Errorf("unable to read remote conf: %v", err)
				continue
			}

			// unmarshal new conf into our runtime conf struct. you can also use channel
			// to implement a signal to notify the system of the changes
			runtime_viper.Unmarshal(&Conf)
		}
	}()
}

// get方法  支持嵌套
type GetViper interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
	AllSettings() map[string]interface{}
}

func getString() {
	viper.GetString("user.name")
	viper.GetString("user.age")
}
