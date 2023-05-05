package main

import (
	"github.com/gin-gonic/gin"
	"goDemo/check"
	"goDemo/model"
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
		errs := check.Check(open)
		if len(errs) == 0 {
			c.JSON(200, model.Response{
				Success: true,
				Msg:     "恭喜您,文件校验通过!",
				ErrMsgs: errs,
			})
		} else {
			c.JSON(200, model.Response{
				Success: false,
				Msg:     "很遗憾,文件还有错误要修改!",
				ErrMsgs: errs,
			})
		}

	})

	r.Run(":7990")
}
