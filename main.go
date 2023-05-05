package main

import (
	"fmt"
	"goDemo/check"
)

func main() {
	errs := check.Check()
	for _, e := range errs {
		fmt.Println(e)
	}
	// 打开 Excel 文件

}
