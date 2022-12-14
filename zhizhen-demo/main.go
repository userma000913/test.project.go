package main

import "fmt"

// 指针  go的指针不能做运算
func main() {
	var a = 2
	// pa的指针指向a的地址
	var pa = &a

	// pa的值改变之后a的值也会跟着改变
	*pa = 3
	fmt.Println(a, *pa, pa == &a)

	a1 := 1
	f1(&a1)
	fmt.Println(a1)

	data := 1
	c := Cache{
		pData: &data,
	}
	// todo
	f2(c)
	fmt.Printf("%d", *c.pData)

}

type Cache struct {
	pData *int
}

func f1(a1 *int) {
	*a1 = 2
}

func f2(c Cache) {
	data := 100
	c.pData = &data
}
