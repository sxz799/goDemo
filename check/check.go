package check

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"goDemo/utils"
	"log"
	"strconv"
	"strings"
)

func Check() (errs []error) {
	excelFile, err := xlsx.OpenFile("/Users/sxz799/Desktop/WinFile/AAA.xlsx")
	if err != nil {
		errs = append(errs, err)
		return
	}
	log.Println("文件读取成功！")
	sheets := excelFile.Sheets
	if len(sheets) != 1 {
		errs = append(errs, errors.New("只能有一个sheet页"))
	}
	rows := sheets[0].Rows

	firstRow := rows[0]

	indexTitleMap := make(map[int]string)

	for i, cell := range firstRow.Cells {
		indexTitleMap[i] = cell.Value
	}

	rows = rows[1:]

	var gsidmp = make(map[string]string)

	for i, row := range rows {
		if row.Cells[0].Value == "" || row.Cells[0].Value == "合计" {
			log.Println("共校验了", i, "行数据")
			break
		}
		capNum := 0
		GSID := ""
		titleValueMap := make(map[string]string)
		// 遍历该行中的所有单元格
		for k, cell := range row.Cells {
			if cell.Value == "" {
				continue
			}
			title := indexTitleMap[k]

			titleValueMap[title] = cell.Value
			//校验资产编码唯一性
			if title == "资产编号" {
				s := gsidmp[cell.Value]
				if s != "" {
					errs = append(errs, errors.New("资产编号"+cell.Value+"已存在"))
				} else {
					gsidmp[cell.Value] = cell.Value
				}
				GSID = cell.Value
			}

			if title == "责任人" || title == "使用人" {
				if strings.Contains(cell.Value, "+") {
					if capNum != len(strings.Split(cell.Value, "+")) {
						errs = append(errs, errors.New("责任人或使用人数量配置异常！资产编号为"+GSID))
					}
				}
			}
			f, ok := utils.TitleCheckFuncMap[title]
			if ok {
				err = f(cell.Value)
				if err != nil {
					errs = append(errs, errors.New(title+"->["+err.Error()+"] 资产编号为："+GSID))
				}
			}

		}

		if titleValueMap["是否计提折旧"] == "是" {
			syyf, _ := strconv.Atoi(titleValueMap["使用月份"])
			ytyf, _ := strconv.Atoi(titleValueMap["已提月份"])
			wtyf, _ := strconv.Atoi(titleValueMap["未计提月份"])
			if syyf != ytyf+wtyf {
				err = errors.New("使用月份不等于已提月份加未提月份,资产编号为：" + titleValueMap["资产编号"])
				break
			}
			float, _ := strconv.ParseFloat(titleValueMap["净残值率(%)"], 64)
			float = float / 100
			if float >= 1 || float < 0 {
				fmt.Println("净产值率不可大于等于1或小于0，资产编号", titleValueMap["资产编号"])
				break
			}
		}

	}
	return
}
