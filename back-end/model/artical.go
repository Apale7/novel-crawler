package model

import "fmt"

type Article struct {
	Directory
	Content string
}

type Directory struct {
	ID    int64  `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type ArticleSlice []Article

func (a ArticleSlice) Len() int {
	return len(a)
}

func (a ArticleSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ArticleSlice) Less(i, j int) bool {
	return a[i].ID < a[j].ID
}

func (a Article) String() string {
	return fmt.Sprintf("%s\n%s\n", a.Title, a.Content)
}
