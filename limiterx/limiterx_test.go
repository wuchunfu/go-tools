package limiterx

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	t.Run("Limiter", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		// 创建一个缓存channel
		var ch = make(chan int)
		go Request(wg, ch)
		// waitGroup在主线程等待所有任务完成
		wg.Wait()
		for num := range ch {
			t.Log("num = ", num)
		}
	})
}

func Request(wg *sync.WaitGroup, ch chan<- int) {
	wg.Add(1)
	defer wg.Done()
	begin := time.Now()
	for i := 1; i <= 12; i++ {
		if !NewLimiter(1*time.Second, 10, "127.0.0.1").Allow() {
			fmt.Println("访问频率过高")
		}
		// 往通道写数据
		ch <- i
	}
	// 不需要再写数据时，关闭channel
	close(ch)
	end := time.Now().Sub(begin).Seconds()
	fmt.Printf("time consumed: %fs\n", end)
}
