package syn_map

import (
	"sync"
	"sync/atomic"
	"time"
)

func Increment(m *sync.Map, g int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := range 1000 {
		m.Store(g*1000+i, i)
	}
}

func synMapTime() (time.Duration, int64) {

	var wg sync.WaitGroup
	var m sync.Map

	start := time.Now()

	for g := range 50 {
		wg.Add(1)
		go Increment(&m, g, &wg)
	}

	wg.Wait()
	lengthTime := time.Since(start)

	// no count, use range
	var cnt atomic.Int64
	m.Range(func(_, _ any) bool {
		cnt.Add(1)
		return true
	})

	return lengthTime, cnt.Load()
}

func Syn_Map() (time.Duration, int64) {
	var sumDur time.Duration
	var sumCnt int64

	for range 10 {
		dur, cnt := synMapTime()
		sumDur += dur
		sumCnt += cnt
	}
	avgDur := sumDur / time.Duration(10)
	avgCnt := sumCnt / int64(10)
	return avgDur, avgCnt
}
