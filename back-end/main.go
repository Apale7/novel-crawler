package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	defer r.Run(":6350")
	collectRouter(r)
}

// func V1() {
// 	var c crawler.Crawler
// 	startTime := time.Now().UnixNano()
// 	c.GetNovel("/38/38662/")
// 	endTime := time.Now().UnixNano()
// 	fmt.Printf("V1总用时: %.4fs", float64(endTime-startTime)/1e9)
// }
