package crawler

import (
	"fmt"
	"novel-crawler-back-end/model"
	"os"
	"sort"
	"strings"

	"sync"

	"github.com/sirupsen/logrus"
)

var (
	novel    string
	articles model.ArticleSlice
	n        int
	tot      int
	lock     sync.Mutex
)

func GetNovel(book string) (err error) {
	var url string
	if !strings.HasPrefix(book, baseURL) {
		url = fmt.Sprintf("%s/%s", baseURL, book)
	} else {
		url = book
	}
	n = 0
	articles = model.ArticleSlice{}
	name, dirs, err := GetList(url)
	if err != nil {
		return
	}

	n = len(dirs)

	var wg sync.WaitGroup
	for _, dir := range dirs {
		wg.Add(1)
		go crawl(dir, &wg)
	}
	wg.Wait()

	//按id排序
	sort.Sort(articles)
	if len(name) == 0 {
		name = "未知小说.txt"
	}

	//写入文件
	w, _ := os.OpenFile("output/"+name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	var novel []string
	for i, a := range articles {
		if i > 0 && a.ID == articles[i-1].ID {
			continue
		}
		novel = append(novel, a.String())
	}
	defer w.Close()
	_, err = w.Write([]byte(strings.Join(novel, "")))
	if err != nil {
		logrus.Errorf("写入文件失败, err: %+v", err)
	}
	return
}

func crawl(dir model.Directory, wg *sync.WaitGroup) {
	content, err := GetContent(dir.URL)
	if err != nil {
		logrus.Errorf("爬取%s失败, err: %+v", dir.Title, err)
		return
	}
	article := model.Article{
		Directory: dir,
		Content:   content,
	}
	lock.Lock()
	articles = append(articles, article)
	lock.Unlock()
	fmt.Printf("已爬取: %s (%d/%d)\n", article.Title, len(articles), n)
	wg.Done()
}
