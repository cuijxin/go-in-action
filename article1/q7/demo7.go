package main

//  使用atomic包来提供对数值类型的安全访问
import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全地对counter加1
		atomic.AddInt64(&counter, 1)
		// 当前goroutine从线程退出，并放回到队列
		runtime.Gosched()
	}
}
