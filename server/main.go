package main

import (
	"github.com/gin-gonic/gin"
	"gsCheck/check"
	"gsCheck/model"
	"strconv"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("dist/index.html")
	r.Static("/dist", "dist")
	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", "")
	})

	r.POST("/api/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		open, _ := file.Open()
		num, errs := check.Check(open)
		extMsg := "[本次共校验" + strconv.Itoa(num) + "行数据]"
		if len(errs) == 0 {
			c.JSON(200, model.Response{
				Success: true,
				Msg:     "恭喜您,文件校验通过!" + extMsg,
				ErrMsgs: errs,
			})
		} else {
			c.JSON(200, model.Response{
				Success: false,
				Msg:     "很遗憾,文件还有" + strconv.Itoa(len(errs)) + "个错误要修改!" + extMsg,
				ErrMsgs: errs,
			})
		}

	})

	r.Run(":7990")
}
