// Increment 是 Container 的一个方法 (method)。
// (c *Container) 部分：叫做方法接收者 (receiver)，类似 Java 里的 `this`。
//   - `c` 就是调用方对象，例如 main 里创建的 Container 实例。
//   - `*Container` 表示接收者是一个指针，保证我们操作的是同一个对象（而不是复制品）。
//
// 为什么要用指针？
//   - Go 默认是值传递，如果写成 (c Container)，每次调用都会拷贝一份 struct，
//     导致锁和 map 都被复制，线程安全彻底失效。
//   - 用 *Container 就和 Java 里的对象引用类似，多个 goroutine 操作的都是同一个实例。
//
// 方法内部：
//   - c.mu.Lock() / defer c.mu.Unlock() → 在整个方法执行期间给 Container 上锁，
//     确保没有其他 goroutine 可以同时写这个 map。
//   - for 循环：写入 1000 个键值对，key 是 g*1000+i，value 是 i。
//   - 相当于 Java 里的 synchronized 方法：
//       synchronized void increment(int g) { ... }
//
// 线程安全设计思路：
//   - Container struct 把数据 (counters) 和锁 (mu) 封装在一起，形成一个线程安全的 map。
//   - 外部代码只需要调用 c.Increment(g)，不用关心加锁问题。
//   - 类似 Java 里的 Collections.synchronizedMap(...) 包装。
//

package map_mutex

import (
	"sync"
	"time"
)

type Container struct {
	mu       sync.Mutex
	mapCount map[int]int
}

func (c *Container) Increment(g int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := range 1000 {
		c.mapCount[g*1000+i] = i
	}
}

// func worker(g int, c *Container, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	c.Increment(g)
// }

func mapMutexTime() (time.Duration, int64) {

	const (
		Total = 50 * 1000
	)

	c := Container{mapCount: make(map[int]int, Total)}

	var wg sync.WaitGroup

	start := time.Now()

	for g := range 50 {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			c.Increment(g)
		}(g)
	}
	wg.Wait()

	lengthTime := time.Since(start)
	count := int64(len(c.mapCount))

	return lengthTime, count
}

func Map_mutex() (time.Duration, int64) {
	var sumDur time.Duration
	var sumCnt int64

	for range 10 {
		d, c := mapMutexTime()
		sumDur += d
		sumCnt += c
	}
	avgDur := sumDur / time.Duration(10)
	avgCnt := sumCnt / int64(10)
	return avgDur, avgCnt
}
