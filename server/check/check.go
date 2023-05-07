package check

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gsCheck/model"
	"gsCheck/utils"
	"io"
	"log"
	"strconv"
	"strings"
)

func Check(r io.Reader) (num int, errs []model.ErrInfo) {
	excelFile, err := excelize.OpenReader(r)
	//excelFile, err := excelize.OpenFile(filename)
	if err != nil {
		errs = append(errs, model.ErrInfo{
			ErrorMsg: err.Error(),
			FixMsg:   "文件读取失败,请按照使用说明导出格式为xlsx的文件进行检测",
		})
		return
	} else {
		fmt.Println("文件读取成功！")
	}

	count := excelFile.SheetCount

	if count != 1 {
		errs = append(errs, model.ErrInfo{
			ErrorMsg: "sheet工作表太多",
			FixMsg:   "请仅保留一个工作表(如果仍提示此错误,请右键点击左下角现在的工作表并点击`取消隐藏工作表`)",
		})
	}
	list := excelFile.GetSheetList()

	cellA3, _ := excelFile.GetCellValue(list[0], "A3")
	cellA3 = strings.ReplaceAll(cellA3, " ", "")
	if !strings.Contains("资产编号,资产名称,资产来源,管理类别,类别名称1,资产状态,是否计提折旧,入账日期,资产原值,累计折旧,折旧方法,资产数量,净残值率(%),净残值,月折旧率(%),月折旧额,年折旧率(%),年折旧额,存放地点,部门名称,责任人,入账时累计折旧,减值准备,已提月份,未计提月份,单位名称,使用部门,使用人,使用月份,计量单位,备注,实际数量,,", cellA3) {
		errs = append(errs, model.ErrInfo{
			ErrorMsg: "表格结构错误",
			FixMsg:   "请将前两行内容清空并合并为1个单元格!",
		})
		return
	}

	rows, _ := excelFile.GetRows(list[0])

	firstRow := rows[2]

	indexTitleMap := make(map[int]string)

	for i, cell := range firstRow {
		indexTitleMap[i] = cell
	}

	rows = rows[3:]

	var gsidmp = make(map[string]string)

	for index, row := range rows {
		if row[0] == "合计" {
			log.Println("共校验了", index, "行数据")
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
				if len(strings.ReplaceAll(cell, " ", "")) < 1 {
					errs = append(errs, model.ErrInfo{
						Line:     index + 4,
						ErrorMsg: "资产编号错误",
						FixMsg:   "资产编号不可为空",
					})
				}
				GSID = cell
				_, ok := gsidmp[GSID]
				if ok {
					errs = append(errs, model.ErrInfo{
						Line:     index + 4,
						ErrorMsg: "资产编号:" + GSID + "已存在",
						FixMsg:   "修改为不重复的编码(比如后面加上一些字母)",
					})
				} else {
					gsidmp[cell] = cell
				}

			}

			if title == "责任人" || title == "使用人" {
				if len(strings.ReplaceAll(cell, " ", "")) < 1 {
					errs = append(errs, model.ErrInfo{
						Line:     index + 4,
						ErrorMsg: "责任人或使用人不可为空！",
						FixMsg:   "填入责任或使用人信息",
					})
				} else if strings.Contains(cell, "+") {
					if capNum != len(strings.Split(cell, "+")) {
						errs = append(errs, model.ErrInfo{
							Line:     index + 4,
							ErrorMsg: "责任人或使用人数量配置异常！",
							FixMsg:   "修改资产数量与使用人的数量(资产数量大于1时,人员要么1个，要么与资产数量相同)",
						})
					}
				}
			}
			f, ok := utils.TitleCheckFuncMap[title]
			if ok {
				correct, errInfo := f(cell)
				if !correct {
					errInfo.Line = index + 4
					errInfo.ErrorMsg = title + errInfo.ErrorMsg
					errs = append(errs, errInfo)
				}
			}

		}

		mkt, ok := titleValueMap["单位名称"]
		if ok {
			correct, errInfo := utils.IsCorrectMKT(mkt)
			if !correct {
				errInfo.Line = index + 4
				errs = append(errs, errInfo)
			} else {
				correct, errInfo = utils.IsCorrectDept(titleValueMap["部门名称"], mkt)
				if !correct {
					errInfo.Line = index + 4
					errs = append(errs, errInfo)
				}
				correct, errInfo = utils.IsCorrectDept(titleValueMap["使用部门"], mkt)
				if !correct {
					errInfo.Line = index + 4
					errs = append(errs, errInfo)
				}

				users, ok := titleValueMap["责任人"]
				if ok {
					if strings.Contains(users, "+") {
						for _, user := range strings.Split(users, "+") {
							correct, errInfo = utils.IsCorrectUser(user, mkt)
							if !correct {
								errInfo.Line = index + 4
								errInfo.ErrorMsg = "责任人" + errInfo.ErrorMsg
								errs = append(errs, errInfo)
							}
						}
					} else {
						correct, errInfo = utils.IsCorrectUser(users, mkt)
						if !correct {
							if !correct {
								errInfo.Line = index + 4
								errInfo.ErrorMsg = "责任人" + errInfo.ErrorMsg
								errs = append(errs, errInfo)
							}
						}
					}

				}

				users2, ok := titleValueMap["使用人"]
				if ok {
					if strings.Contains(users2, "+") {
						for _, user := range strings.Split(users2, "+") {
							correct, errInfo = utils.IsCorrectUser(user, mkt)
							if !correct {
								if !correct {
									errInfo.Line = index + 4
									errInfo.ErrorMsg = "使用人" + errInfo.ErrorMsg
									errs = append(errs, errInfo)
								}
							}
						}
					} else {
						correct, errInfo = utils.IsCorrectUser(users2, mkt)
						if !correct {
							if !correct {
								errInfo.Line = index + 4
								errInfo.ErrorMsg = "使用人" + errInfo.ErrorMsg
								errs = append(errs, errInfo)
							}
						}
					}

				}
			}

		} else {
			errs = append(errs, model.ErrInfo{
				Line:     index + 4,
				ErrorMsg: "单位名称不可为空",
				FixMsg:   "填入提供的组织架构中的门店名称",
			})
		}

		if titleValueMap["是否计提折旧"] == "是" {
			syyf, _ := strconv.Atoi(titleValueMap["使用月份"])
			ytyf, _ := strconv.Atoi(titleValueMap["已提月份"])
			wtyf, _ := strconv.Atoi(titleValueMap["未计提月份"])
			if syyf != ytyf+wtyf {
				errs = append(errs, model.ErrInfo{
					Line:     index + 4,
					ErrorMsg: "使用月份不等于已提月份加未提月份",
					FixMsg:   "校对使用月份，已提月份，未提月份",
				})

			}
			float, _ := strconv.ParseFloat(titleValueMap["净残值率(%)"], 64)
			float = float / 100
			if float >= 1 || float < 0 {

				errs = append(errs, model.ErrInfo{
					Line:     index + 4,
					ErrorMsg: "净产值率错误",
					FixMsg:   "净产值率不可大于等于1或小于0",
				})
			}
		}

	}
	log.Println("校验完毕")
	num = len(rows)
	return
}
