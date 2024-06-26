package base

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.WaitGroup提供一个可等待协程完成的group,可用于等待一组协程的完成同时可配合channel实现限流操作

var wg = &sync.WaitGroup{}

func Test_WaitGroup(t *testing.T) {
	channel := make(chan int, 5)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			channel <- i
			fmt.Printf("channel value is %v\n", i)
			if len(channel) == 5 {
				time.Sleep(3*time.Second)
			}
			defer func() {
				<- channel
				wg.Done()
			}()
		}(i)
	}

	wg.Wait()
}