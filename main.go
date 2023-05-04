package main

import (
	"fmt"
	"goDemo/utils"
	"log"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

var gsidmp = make(map[string]string)

func main() {
	// 打开 Excel 文件
	excelFile, err := xlsx.OpenFile("/Users/sxz799/Desktop/testfile.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("文件读取成功！")
	sheets := excelFile.Sheets
	if len(sheets) != 1 {
		log.Println("只能有一个sheet页面")
		return
	}
	rows := sheets[0].Rows

	firstRow := rows[0]

	titleMap := make(map[int]string)

	for i, cell := range firstRow.Cells {
		titleMap[i] = cell.Value
	}

	rows = rows[1:]

	for i, row := range rows {
		if row.Cells[0].Value == "" {
			log.Println("共校验了", i, "行数据")
			break
		}
		capnum := 0
		gsid := ""
		// 遍历该行中的所有单元格
		for k, cell := range row.Cells {
			title := titleMap[k]
			//校验资产编码唯一性
			if title == "资产编码" {
				s := gsidmp[cell.Value]
				if s != "" {
					log.Println("资产编码已存在！", cell.Value)
					break
				} else {
					gsidmp[cell.Value] = cell.Value
				}
				gsid = cell.Value
			}

			if title == "资产数量" {
				capnum, err = strconv.Atoi(cell.Value)
				if err != nil {
					log.Println("资产数量异常！资产编码为", gsid)
					break
				}
			}

			if title == "责任人" || title == "使用人" {
				if strings.Contains(cell.Value, "+") {
					if capnum != len(strings.Split(cell.Value, "+")) {
						log.Println("责任人或使用人配置异常！资产编码为", gsid)
					}
				}
			}

			f := utils.TitleFunc[title]
			b, err := f(cell.Value)
			if !b {
				fmt.Println(err.Error())
			}

		}

	}
}
