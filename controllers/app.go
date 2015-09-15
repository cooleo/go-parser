package controllers

import (
	"fmt"
	"github.com/herotl2005/go-parser/models"
	"github.com/herotl2005/go-parser/parser/nct"
	"github.com/herotl2005/go-parser/parser/tvzing"
    "github.com/herotl2005/go-parser/parser/clipvn"
    elastigo "github.com/mattbaird/elastigo/lib"
)

type App struct {
	//*revel.Controller
}

func DoParser(dao *models.Dao, c *elastigo.Conn) {

	pages := dao.FindParsePages()
	for i, _ := range pages {

		//var page *models.PageHtml
		page := pages[i]

		fmt.Printf("page:%s\n", page.Url)
		if page.HostName == "v.nhaccuatui.com" {
			nct.Print(dao, page, c)
		}
		if page.HostName == "tv.zing.vn" {
			tvzing.Print(dao, page, c)
		}
        if page.HostName == "clip.vn" {
			clipvn.Print(dao, page, c)
		}
		page.Html = ""
		page.Parsed = true

		result:= dao.UpdatePage(page)
		fmt.Println("result:", result)
	}
}
