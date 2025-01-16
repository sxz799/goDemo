package utils

import (
	"database/sql"
	"fmt"
	"github.com/sijms/go-ora/v2"
	"gsCheck/check"
	"gsCheck/model"
	"strings"
)

func InitCheckFuncMap() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("[旧]门店数量：%d\n", len(check.MktMap)))
	s.WriteString(fmt.Sprintf("[旧]部门数量：%d\n", len(check.OrgMap)))
	s.WriteString(fmt.Sprintf("[旧]员工数量：%d\n", len(check.UserMap)))

	buildUrl := go_ora.BuildUrl("10.160.55.112", 1521, "csrac", "capital", "capital", nil)
	conn, err := sql.Open("oracle", buildUrl)
	if err != nil {
		fmt.Println("ERR1,链接数据库失败")
		panic(err)
	}
	err2 := conn.Ping()
	if err2 != nil {
		fmt.Println("ERR2,测试数据库失败")
		panic(err2)
	}

	sqlStr1 := `
	SELECT
	   d2.DEPT_NAME mkt,
		d1.DEPT_NAME dept
	
	FROM
		SYS_DEPT d1
		LEFT JOIN SYS_DEPT d2 on d1.MKT_CODE=d2.DEPT_CODE
	WHERE
		d1.IS_END = '1'
	`
	orgRows, err3 := conn.Query(sqlStr1)
	defer orgRows.Close()
	if err3 != nil {
		fmt.Println("查询组织架构失败")
		panic(err3)
	}
	check.OrgMap = make(map[model.Organization]struct{})
	check.MktMap = make(map[string]struct{})
	for orgRows.Next() {
		var org model.Organization
		orgRows.Scan(&org.Dept, &org.Mkt)
		check.MktMap[org.Mkt] = struct{}{}
		check.OrgMap[org] = struct{}{}
		fmt.Println(org.Mkt)
	}

	sqlStr2 := `
SELECT
	u.REALNAME name,
	d.DEPT_NAME mkt 
FROM
	SYS_USER_MANGE_MKT um
	LEFT JOIN SYS_USER u on u.id=um.USER_ID
	LEFT JOIN SYS_DEPT d ON um.mkt = d.DEPT_CODE

`
	userRows, err3 := conn.Query(sqlStr2)
	defer userRows.Close()
	if err3 != nil {
		fmt.Println("查询用户信息失败")
		panic(err3)
	}
	check.UserMap = make(map[model.User]struct{})
	for userRows.Next() {
		var u model.User
		userRows.Scan(&u.Name, &u.Mkt)
		check.UserMap[u] = struct{}{}
		fmt.Println(u.Name,u.Mkt)
	}

	fmt.Println("数据初始化完成！")

	s.WriteString(fmt.Sprintf("门店数量：%d\n", len(check.MktMap)))
	s.WriteString(fmt.Sprintf("部门数量：%d\n", len(check.OrgMap)))
	s.WriteString(fmt.Sprintf("员工数量：%d\n", len(check.UserMap)))
	return s.String()
}
