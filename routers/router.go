package routers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"myblob/pkg/setting"
	"myblob/pkg/utils"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.SetFuncMap(template.FuncMap{
		"md5":utils.Md5,
	})

	r.LoadHTMLGlob("views/*/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			gin.H{
			"title": "aibyte",
			"data": struct {
				Title string `json:"title"`
				Content string `json:"content"`
			}{
				Title: "aibyte",
				Content: "首页",
			},
		},
		)
	})
	return r
}