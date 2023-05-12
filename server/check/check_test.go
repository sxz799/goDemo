package check

import (
	"gsCheck/model"
	"log"
	"os"
	"testing"
)

func TestPreCheck(t *testing.T) {

	// 测试 xlsx 文件
	errs := make([]model.ErrInfo, 0)
	open, err := os.Open("/Users/sxz799/Desktop/WinFile/淄博区域资产导入表/淄博店/淄博区域商城店固定资产5.9.xlsx")
	if err != nil {
		log.Println("文件打开失败")
		return
	}
	_, errs = PreCheck("xlsx", open)
	for _, e := range errs {
		log.Println(e)
	}

	// 测试 xls 文件

}
