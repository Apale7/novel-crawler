package crawler

import (
	"fmt"
	"novel-crawler-back-end/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

var (
	name string
)

// 返回不含content字段的article数组
func GetList(url string) (name string, dir []model.Directory, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		logrus.Errorf("获取章节列表失败, err: %+v", err)
		return
	}
	name = doc.Find("#info").Find("h1").Text() + ".txt"
	doc.Find("#list").Children().Find("a").Each(func(i int, this *goquery.Selection) {
		title := this.Text()
		title = strings.ReplaceAll(title, "?", " ")
		href, _ := this.Attr("href")

		d := model.Directory{ID: getID(href), Title: title, URL: baseURL + href}
		fmt.Println(d)
		dir = append(dir, d)
	})
	fmt.Println(len(dir))
	return
}

func getID(href string) int64 {
	attr := strings.Split(href, "/")
	tmp := strings.Replace(attr[len(attr)-1], ".html", "", 1)
	id, _ := strconv.ParseInt(tmp, 10, 64)
	return id
}
