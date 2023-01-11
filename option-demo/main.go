package main

import "fmt"

// 可变参数option
func main() {
	NewFriend2(
		"看电影",
		WithSex(1),
		WithAge(20),
	)
}

// NewFriend 寻找志同道合朋友
// sex和age是非必传参数
func NewFriend(sex int, age int, hobby string) (string, error) {

	// 逻辑处理 ...

	return "", nil
}

// option可变参数实现方式

type Option func(*option)
type option struct {
	sex int
	age int
}

func WithSex(sex int) Option {
	return func(o *option) {
		o.sex = sex
	}
}

func WithAge(age int) Option {
	return func(o *option) {
		o.age = age
	}
}

func NewFriend2(hobby string, opts ...Option) (string, error) {
	opt := new(option)
	for _, f := range opts {
		f(opt)
	}
	fmt.Println(opt.age)
	return "", nil
}
