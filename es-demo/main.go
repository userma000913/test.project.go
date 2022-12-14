package main

import (
	"github.com/olivere/elastic/v7"
	"log"
)

var ES *elastic.Client
var err error

func initEs() {

	//默认不记录日志

	// 本地启动时，需要设置为false，具体原因参考 https://juejin.cn/post/6895371414179446798
	sniffOpt := elastic.SetSniff(false)
	ES, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"), sniffOpt,
		elastic.SetTraceLog(new(TraceLog)))
	//ES, err = elastic.NewClient(elastic.SetURL(host))
	// 打印查询ES的语句
	//ES, err = elastic.NewClient(elastic.SetURL(host), elastic.SetTraceLog(new(TraceLog)))

	if err != nil {
		log.Println("ElasticSearch连接失败。。。", err)
	} else {
		log.Println("ElasticSearch已连接 >>> ")
	}
}

// TraceLog 实现 elastic.Logger 接口
type TraceLog struct{}

// Printf 实现输出ES日志
func (TraceLog) Printf(format string, v ...interface{}) {
	log.Println("ES查询日志：", v)
}

func main() {
	initEs()

}

func Es() {

}
