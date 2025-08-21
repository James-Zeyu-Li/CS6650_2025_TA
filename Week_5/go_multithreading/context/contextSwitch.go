package context

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const RoundTrips = 1000000

//启动 G1 和 G2。Go 调度器很有可能会把它们分配到两个不同的 OS 线程上，
// 比如 G1 在线程 T1 上，G2 在线程 T2 上。这两个线程可以并行地运行在两个不同的 CPU 核心上。

func runTest(maxProcess, roundTrips int) (timeCost time.Duration, avg time.Duration) {

	old := runtime.GOMAXPROCS(maxProcess) //limit to only one thread runs
	defer runtime.GOMAXPROCS(old)         // return everything to it's original, clean up -》 why defer
	// if don't fix it 后续的所有 Goroutine：之后创建的任何新的 Goroutine 都会受到这个限制。
	// 已存在的其他 Goroutine：整个 Go 程序的运行时调度器（Runtime Scheduler）都会被切换到新的模式下运行。

	ch := make(chan struct{}) //channel no buffer

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		var token struct{}
		for i := 0; i < RoundTrips; i++ {
			ch <- token //blocks the channel until <- ch
			<-ch        // clocks until another channel do ch<-token
		}
	}()

	go func() {
		defer wg.Done()
		var token struct{}
		for i := 0; i < RoundTrips; i++ {
			<-ch        // A -> B
			ch <- token // B -> A
		}
	}()

	start := time.Now()

	wg.Wait()

	timeCost = time.Since(start)

	avg = timeCost / time.Duration(2*roundTrips)
	return

}

func ContextSwitchTest() {

	total1, avg1 := runTest(1, RoundTrips)
	fmt.Printf("Single-thread   total: %v, avg switch: %v\n", total1, avg1)

	total2, avg2 := runTest(runtime.NumCPU(), RoundTrips)
	fmt.Printf("Multi-thread    total: %v, avg switch: %v\n", total2, avg2)
}
