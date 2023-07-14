package crawl

import (
	"fmt"
	"log"
)

// 令牌是一个计数信号量
// 确保并发请求限制在20个以内
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取令牌
	list, err := Extract(url)
	<-tokens // 释放令牌

	if err != nil {
		log.Print(err)
	}
	return list
}
