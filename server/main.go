package main

import (
	"github.com/gin-gonic/gin"
	"gsCheck/check"
	"gsCheck/model"
	"os"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()

	_, err := os.Stat("dist")
	if err == nil {
		r.LoadHTMLGlob("dist/index.html")
		r.Static("/dist", "dist")
		r.GET("/", func(context *gin.Context) {
			context.HTML(200, "index.html", "")
		})
	}

	r.POST("/api/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		open, _ := file.Open()
		before := time.Now()

		fileType := ""

		filename := file.Filename
		switch {
		case filename[len(filename)-4:] == ".xls":
			fileType = "xls"
		case filename[len(filename)-4:] == "xlsx":
			fileType = "xlsx"
		}
		num, errs := check.PreCheck(fileType, open)
		after := time.Now()
		duration := after.Sub(before)

		extMsg := " [本次共校验" + strconv.Itoa(num) + "行数据,共计耗时" + strconv.FormatInt(duration.Milliseconds(), 10) + "ms]"
		if len(errs) == 0 {
			c.JSON(200, model.Response{
				Success:  true,
				Msg:      "恭喜您,文件校验通过!" + extMsg,
				ErrInfos: errs,
			})
		} else {

			switch {
			case len(errs) >= 20:
				c.JSON(200, model.Response{
					Success:  false,
					Msg:      "很遗憾,文件还有" + strconv.Itoa(len(errs)) + "个错误要修改-_-!" + extMsg,
					ErrInfos: errs,
				})
			case len(errs) < 20 && len(errs) > 10:
				c.JSON(200, model.Response{
					Success:  false,
					Msg:      "努努力,就还剩" + strconv.Itoa(len(errs)) + "个错误了!" + extMsg,
					ErrInfos: errs,
				})
			case len(errs) < 10:
				c.JSON(200, model.Response{
					Success:  false,
					Msg:      "加把劲,还有最后" + strconv.Itoa(len(errs)) + "个错误了!" + extMsg,
					ErrInfos: errs,
				})
			}

		}

	})

	r.Run(":7990")
}
