package map_routine

import (
	"fmt"
	"sync"
)

func goroutine(m map[int]int, g int, wg *sync.WaitGroup) {
	defer wg.Done() // notify go routine done
	for i := range 1000 {
		m[g*1000+i] = i
	}
}

func Map_routine() {

	m := make(map[int]int)

	var wg sync.WaitGroup

	for g := range 50 {
		wg.Add(1)
		go goroutine(m, g, &wg)
	}

	wg.Wait()
	fmt.Println("len(m) = ", len(m))
}
