package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// rune and string
func main() {
	s := "hello 你好ok世界"
	fmt.Println(len(s))
	fmt.Println("rune size is ", utf8.RuneCount([]byte(s)))

	bytes := []byte(s)
	for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", r)
	}
	fmt.Println()
	res := ""
	for _, c := range []rune(s) {

		res += fmt.Sprintf("%c", c)
	}
	fmt.Println(res)
	println("------------------")

	s1 := "我 是  你  爸爸"
	fmt.Println(strings.Split(s1, " "), len(strings.Split(s1, " ")))
	// Fields专门切割空格，多个空格的也可以操作
	fmt.Println(strings.Fields(s1), len(strings.Fields(s1)))

	//for _, r := range []rune(s1) {
	//	fmt.Printf("%s",string(r))
	//}

	println("------------------")
	f()
}

func f() {
	s := "{\"msg_type\":\"interactive\",\"card\":{\"header\":{\"template\":\"green\",\"title\":{\"content\":\"%s\",\"tag\":\"plain_text\"}},\"elements\":[{\"fields\":[{\"is_short\":true,\"text\":{\"content\":\"**时间**\\n%s\",\"tag\":\"lark_md\"}},{\"is_short\":true,\"text\":{\"content\":\"**项目**\\n%s\",\"tag\":\"lark_md\"}}],\"tag\":\"div\"},{\"tag\":\"div\",\"text\":{\"content\":\"<at id=all></at>\",\"tag\":\"lark_md\"}},{\"tag\":\"hr\"},{\"tag\":\"div\",\"text\":{\"content\":%#v,\"tag\":\"lark_md\"}}]}}"
	fmt.Println(len(s))
	fmt.Println(len([]rune(s)))

	r := make([]string, 0, 10000)

	rt := []rune(s)
	res := ""
	for i := range rt {
		x := string(rt[i])
		fmt.Println("第", i+1, "个字符为:", x)
		res += x
		r = append(r, x)

	}
	//fmt.Println(res)
	//fmt.Println(strings.Contains(res, "elment"))
	//fmt.Printf("%+v", res)
	join := strings.Join(r, "")
	fmt.Printf("%s", join)
}
