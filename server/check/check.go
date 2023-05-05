package check

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"goDemo/model"
	"goDemo/utils"
	"io"
	"log"
	"strconv"
	"strings"
)

func Check(r io.Reader) (errs []model.ErrMsg) {
	excelFile, err := excelize.OpenReader(r)
	//excelFile, err := excelize.OpenFile(filename)
	if err != nil {
		errs = append(errs, model.ErrMsg{
			Msg: err.Error(),
		})
		return
	} else {
		fmt.Println("文件读取成功！")
	}

	count := excelFile.SheetCount
	log.Println(count)
	if count != 1 {
		errs = append(errs, model.ErrMsg{
			Msg: "请仅保留一个工作表(如果仍提示此错误,请右键点击左下角现在的工作表并点击`取消隐藏工作表`)",
		})
	}
	list := excelFile.GetSheetList()
	for _, l := range list {
		fmt.Println(l)
	}
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
					errs = append(errs, model.ErrMsg{
						Msg: "[资产编号:" + GSID + "]" + "资产编号已存在",
					})
				} else {
					gsidmp[cell] = cell
				}
				GSID = cell
			}

			if title == "责任人" || title == "使用人" {
				if strings.Contains(cell, "+") {
					if capNum != len(strings.Split(cell, "+")) {
						errs = append(errs, model.ErrMsg{
							Msg: "[资产编号:" + GSID + "]" + "责任人或使用人数量配置异常！",
						})
					}
				}
			}
			f, ok := utils.TitleCheckFuncMap[title]
			if ok {
				err = f(cell)
				if err != nil {
					errs = append(errs, model.ErrMsg{
						Msg: "[资产编号:" + GSID + "]" + title + err.Error(),
					})
				}
			}

		}

		if titleValueMap["是否计提折旧"] == "是" {
			syyf, _ := strconv.Atoi(titleValueMap["使用月份"])
			ytyf, _ := strconv.Atoi(titleValueMap["已提月份"])
			wtyf, _ := strconv.Atoi(titleValueMap["未计提月份"])
			if syyf != ytyf+wtyf {
				errs = append(errs, model.ErrMsg{
					Msg: "[资产编号:" + GSID + "]" + "使用月份不等于已提月份加未提月份",
				})

			}
			float, _ := strconv.ParseFloat(titleValueMap["净残值率(%)"], 64)
			float = float / 100
			if float >= 1 || float < 0 {
				errs = append(errs, model.ErrMsg{
					Msg: "[资产编号:" + GSID + "]" + "净产值率不可大于等于1或小于0",
				})
			}
		}

	}
	return
}
