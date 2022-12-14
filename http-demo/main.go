package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("", "") // 设置请求头
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 重定向检查
			fmt.Printf("重定向检查:%s\n", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response)
}
