package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) setValue(v int) {
	if t == nil {
		fmt.Println("tree is nil")
		return
	}
	t.value = v
}

func main() {
	//var r *tree // r == nil
	var r1 tree // 这里的r1就不是nil，而是空的tree结构,也就是结构体零值，其对应的指针类型才是nil
	fmt.Println(r1)
	// 这里想说明，r==nil的情况下也可以调用方法，不像java，直接NPE
	r1.setValue(1)
	// 类型转换一下
	t := tr(r1)
	println(t.getValue())
	tt := tree2{
		t: &r1,
	}
	tt.print()

}

type tr tree

type tree2 struct {
	t *tree
}

func (t2 *tree2) print() {
	fmt.Println(t2.t)
}
func (t *tr) getValue() int {
	return t.value
}
