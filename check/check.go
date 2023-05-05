package check

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"goDemo/utils"
	"log"
	"strconv"
	"strings"
)

func Check() (errs []error) {
	excelFile, err := excelize.OpenFile("/Users/sxz799/Desktop/济宁兖州店导入模板-固定资产.xlsx")
	if err != nil {

		errs = append(errs, err)
		return
	}
	fmt.Println("文件读取成功！")
	count := excelFile.SheetCount

	if count != 1 {
		errs = append(errs, errors.New("只能有一个sheet页"))
	}
	list := excelFile.GetSheetList()

	rows, _ := excelFile.GetRows(list[0])

	firstRow := rows[0]

	indexTitleMap := make(map[int]string)

	for i, cell := range firstRow {
		indexTitleMap[i] = cell
	}

	rows = rows[1:]

	var gsidmp = make(map[string]string)

	for i, row := range rows {
		if row[0] == "" || row[0] == "合计" {
			log.Println("共校验了", i, "行数据")
			break
		}
		capNum := 0
		GSID := ""
		titleValueMap := make(map[string]string)
		// 遍历该行中的所有单元格
		for k, cell := range row {
			title := indexTitleMap[k]

			titleValueMap[title] = cell
			//校验资产编码唯一性
			if title == "资产编号" {
				s := gsidmp[cell]
				if s != "" {
					errs = append(errs, errors.New("[资产编号:"+GSID+"]"+"资产编号已存在"))
				} else {
					gsidmp[cell] = cell
				}
				GSID = cell
			}

			if title == "责任人" || title == "使用人" {
				if strings.Contains(cell, "+") {
					if capNum != len(strings.Split(cell, "+")) {
						errs = append(errs, errors.New("[资产编号:"+GSID+"]"+"责任人或使用人数量配置异常！"))
					}
				}
			}
			f, ok := utils.TitleCheckFuncMap[title]
			if ok {
				err = f(cell)
				if err != nil {
					errs = append(errs, errors.New("[资产编号:"+GSID+"]"+title+err.Error()))
				}
			}

		}

		if titleValueMap["是否计提折旧"] == "是" {
			syyf, _ := strconv.Atoi(titleValueMap["使用月份"])
			ytyf, _ := strconv.Atoi(titleValueMap["已提月份"])
			wtyf, _ := strconv.Atoi(titleValueMap["未计提月份"])
			if syyf != ytyf+wtyf {
				errs = append(errs, errors.New("[资产编号:"+GSID+"]"+"使用月份不等于已提月份加未提月份"))
				break
			}
			float, _ := strconv.ParseFloat(titleValueMap["净残值率(%)"], 64)
			float = float / 100
			if float >= 1 || float < 0 {
				errs = append(errs, errors.New("[资产编号:"+GSID+"]"+"净产值率不可大于等于1或小于0"))
				break
			}
		}

	}
	return
}
