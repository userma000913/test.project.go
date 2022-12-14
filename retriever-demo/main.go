package main

import (
	"fmt"
)

type retriever interface {
	get(url string) string
}

type retrieverImpl struct {
}
type retrieverImpl2 struct {
}

// 这样是指针类型
func (r *retrieverImpl2) get(url string) string {
	return url
}

func (r retrieverImpl) get(url string) string {
	return url
}

func show(r retriever) string {
	return r.get("test")
}

func main() {
	var r retriever

	r = retrieverImpl{}
	println(show(r))

	// 判断类型方法
	// 1.switch .(type)
	switch v := r.(type) {
	case retrieverImpl:
		fmt.Println(v.get("test1"))
	case *retrieverImpl2:
		fmt.Println(v.get("test2"))
	}
	var r1 retriever
	r1 = &retrieverImpl2{}
	println(show(r1))

	switch v := r1.(type) {
	case retrieverImpl:
		fmt.Println(v.get("test1"))
	case *retrieverImpl2:
		fmt.Println(v.get("test2"))
	}

	// 2.类型断言

}

type A interface {
	a()
}

type B interface {
	b()
}

// 接口组合 也是继承的意思
type AAndB interface {
	A
	B
}

type Impl struct {
}

func (i Impl) a() {

}
func (i Impl) b() {

}

func f2() {
	var ab AAndB
	ab = Impl{}
	fmt.Println(ab)
}
