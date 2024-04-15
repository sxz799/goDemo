package check

import "gsCheck/model"

// TitleCheckFuncMap 标题和对应的校验函数关系
var TitleCheckFuncMap = make(map[string]func(str string) (bool, model.ErrInfo))

func init() {
	TitleCheckFuncMap["资产名称"] = IsCorrectName
	TitleCheckFuncMap["资产来源"] = IsCorrectComeFrom
	TitleCheckFuncMap["管理类别名称"] = IsCorrectManageType
	TitleCheckFuncMap["资产状态名称"] = IsCorrectStatus
	TitleCheckFuncMap["是否计提折旧"] = IsCorrectJiTi
	TitleCheckFuncMap["资产原值"] = IsDoubleNum
	TitleCheckFuncMap["累计折旧"] = IsDoubleNum
	TitleCheckFuncMap["折旧方法名称"] = IsCorrectZJSF
	TitleCheckFuncMap["资产数量"] = IsIntNum
	TitleCheckFuncMap["净残值率(%)"] = IsCorrectRate
	TitleCheckFuncMap["净残值"] = IsDoubleNum
	TitleCheckFuncMap["月折旧率(%)"] = IsCorrectRate
	TitleCheckFuncMap["月折旧额"] = IsDoubleNum
	TitleCheckFuncMap["年折旧率(%)"] = IsCorrectRate
	TitleCheckFuncMap["年折旧额"] = IsDoubleNum
	TitleCheckFuncMap["入账时累计折旧"] = IsDoubleNum
	TitleCheckFuncMap["减值准备"] = IsDoubleNum
	TitleCheckFuncMap["已提月份"] = IsIntNum
	TitleCheckFuncMap["剩余月份"] = IsIntNum
	// TitleCheckFuncMap["入账日期"] = IsCorrectBuyDate
	//TitleCheckFuncMap["单位名称"] = IsCorrectMKT
	//TitleCheckFuncMap["使用部门"] = IsCorrectDept
	//TitleCheckFuncMap["使用人"] = IsCorrectUser
	//TitleCheckFuncMap["部门名称"] = IsCorrectDept
	//TitleCheckFuncMap["责任人"] = IsCorrectUser
	TitleCheckFuncMap["使用月份"] = IsIntNum
	TitleCheckFuncMap["存放地点名称"] = IsCorrectPlace
	TitleCheckFuncMap["备注"] = IsCorrectMemo
	TitleCheckFuncMap["实际数量"] = IsIntNum
}
