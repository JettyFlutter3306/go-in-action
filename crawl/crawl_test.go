package crawl

import (
	"fmt"
	"testing"
)

func Test_crawl(t *testing.T) {
	workList := make(chan []string)  // 可能重复的URL列表
	unseenLinks := make(chan string) // 去重后的URL列表

	args1 := []string{
		"http://gopl.io",
		"https://golang.org/help",
	}

	args2 := []string{
		"https://golang.org/doc",
		"https://golang.org/blog",
	}

	go func() {
		workList <- args1
		workList <- args2
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	// 主goroutine对URL列表进行去重
	// 把并没有爬去过的条目发送给爬虫程序
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

	fmt.Println(workList)
}
