package crawler

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://www.xxxbiquge.com"
)

type Crawler struct {
}

func analyzeHTML(url string) (content string, err error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	content = doc.Find("#content").Text()
	re := regexp.MustCompile("\\s+")
	content = re.ReplaceAllString(content, "\n")
	re = regexp.MustCompile("^\n")

	content = re.ReplaceAllString(content, "")
	content = strings.ReplaceAll(content, "    天才一秒记住本站地址：[新笔趣阁]\nhttp://www.xxxbiquge.com/最快更新！无广告！", "")
	content = strings.ReplaceAll(content, "章节错误,点此报送(免注册),\n报送后维护人员会在两分钟内校正章节内容,请耐心等待。", "")
	return content, nil
}

func GetContent(url string) (content string, err error) {
	content, err = analyzeHTML(url)
	return
}
