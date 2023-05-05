package main

import (
	"goDemo/check"
	"log"
)

func main() {
	errs := check.Check()
	for _, e := range errs {
		log.Println(e)
	}
	// 打开 Excel 文件

}
