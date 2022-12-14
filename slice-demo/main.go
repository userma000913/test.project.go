package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1) // 正常
	fmt.Println(s2) // 在s1的基础上，去切割原始数组  也就是说 s1在view的时候，是从3开始的，
	// 但是结束是到原始数组末尾，只是由于切片尾部控制，不显示也不能拿出来而已

}
