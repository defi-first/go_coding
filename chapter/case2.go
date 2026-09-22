// 多个goroutine，打印出升序的数字，排他锁实现
package chapter

import (
	"fmt"
	"sync"
)

func Case2() {
	wg := sync.WaitGroup{}
	count := 0
	limit := 100
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.Lock()
				if count >= limit {
					mu.Unlock()
					return
				}
				count++
				fmt.Println(count)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
}
