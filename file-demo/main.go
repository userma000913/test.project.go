package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("/Users/amber/mygo/git.smallma.com/test.project/file-demo/abc.txt")
	f(file)

	s := `
   abc 
	def

p
`
	// 将字符串转为Reader
	f(strings.NewReader(s))

}

func f(reader io.Reader) {
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}
