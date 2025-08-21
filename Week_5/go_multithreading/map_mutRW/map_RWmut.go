package map_mutRW

import (
	"sync"
	"time"
)

type Container struct {
	rwmu     sync.RWMutex
	mapCount map[int]int
}

func (c *Container) Increment(g int) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()

	for i := range 1000 {
		c.mapCount[g*1000+i] = i
	}
}

// func (c *Container) Get( key int)(int, bool){
// 	c.rwmu.RLock()
// 	defer c.rwmu.RUnlock()

// 	v, ok := c.mapCount[key]
// 	return v, ok
// }

// func (c *Container) Length() int {
// 	c.rwmu.RLock()
// 	defer c.rwmu.RUnlock()
// 	return len(c.mapCount)
// }

func mapRWmuTime() (time.Duration, int64) {
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

func Map_RW() (time.Duration, int64) {
	var sumDur time.Duration
	var sumCnt int64

	for range 10 {
		d, c := mapRWmuTime()
		sumDur += d
		sumCnt += c
	}
	avgDur := sumDur / time.Duration(10)
	avgCnt := sumCnt / int64(10)
	return avgDur, avgCnt
}
