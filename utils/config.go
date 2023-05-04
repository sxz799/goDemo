package utils

// TitleFunc 标题和对应的校验函数关系
var TitleFunc = make(map[string]func(str string) (bool, error))

func init() {
	TitleFunc["资产来源"] = IsCorrectComeFrom
	//...
}
