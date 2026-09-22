// 多个goroutine，打印出升序的数字，chan实现
package chapter

import (
	"fmt"
	"sync"
)

func Case1() {
	wg := sync.WaitGroup{}
	limit := 100
	ch := make(chan struct{}, 1)
	count := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				ch <- struct{}{}
				if count >= limit {
					<-ch
					return
				}
				count++
				fmt.Println(count)
				<-ch
			}
		}()
	}

	wg.Wait()
	close(ch)
}
