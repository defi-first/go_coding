// 交替打印数字和字母，字母优先
package chapter

import (
	"fmt"
	"sync"
)

func Case3() {
	length := 52

	letterCount := length / 2
	numberCount := length/2 - 1

	letterTurn := make(chan struct{})
	numberTurn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	// 打印字母
	go func() {
		defer wg.Done()

		for i := 0; i < letterCount; i++ {
			<-letterTurn

			fmt.Println(string(rune('A' + i)))

			// 还有数字时，通知数字协程
			if i < numberCount {
				numberTurn <- struct{}{}
			}
		}
	}()

	// 打印数字
	go func() {
		defer wg.Done()

		for i := 0; i < numberCount; i++ {
			<-numberTurn

			fmt.Println(i % 10)

			// 通知字母协程继续打印
			letterTurn <- struct{}{}
		}
	}()

	// 先唤醒字母协程，保证字母优先
	letterTurn <- struct{}{}

	wg.Wait()
}
