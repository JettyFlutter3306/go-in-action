package concurrent

import (
	"context"
	"fmt"
	"sync"
)

func ContextValue() {
	wg := sync.WaitGroup{}
	ctx1 := context.WithValue(context.Background(), "title", "Go Go Go 1")
	ctx2 := context.WithValue(ctx1, "title2", "Go Go Go 2")
	ctx3 := context.WithValue(ctx2, "title3", "Go Go Go 3")

	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		if v := c.Value("title"); v != nil {
			fmt.Println("Eureka: ", v)
		} else {
			fmt.Println("Not Found: title")
		}
	}(ctx3)

	wg.Wait()
}
