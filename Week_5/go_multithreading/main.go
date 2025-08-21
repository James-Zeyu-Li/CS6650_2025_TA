package main

import ( // "github.com/James-Zeyu-Li/go_multithreading/counterAtomic"
	// "github.com/James-Zeyu-Li/go_multithreading/map_routine"
	// "github.com/James-Zeyu-Li/go_multithreading/map_mutRW"
	// "github.com/James-Zeyu-Li/go_multithreading/map_mutex"
	// "github.com/James-Zeyu-Li/go_multithreading/syn_map"
	//File IO
	// "github.com/James-Zeyu-Li/go_multithreading/doc_write"
	// Context Switch
	"github.com/James-Zeyu-Li/go_multithreading/context"
)

func main() {
	// counterAtomic.AtomicT()
	// map_routine.Map_routine()
	// map_mutex.Map_mutex()
	// map_mutRW.Map_RWmu()
	// syn_map.Map_mutex()

	// avgMutex, cnt := map_mutex.Map_mutex()
	// fmt.Printf("Mutex len(m) = %d, Time = %s (~%d ms)\n", cnt, avgMutex, avgMutex.Milliseconds())

	// avgRW, cnt := map_mutRW.Map_RW()
	// fmt.Printf("Mutex len(m) = %d, Time = %s (~%d ms)\n", cnt, avgRW, avgRW.Milliseconds())

	// synMap, cnt := syn_map.Syn_Map()
	// fmt.Printf("sync.Map size=%d time=%s (~%d ms)\n", cnt, synMap, synMap.Milliseconds())

	// doc_write.WriteTime()

	context.ContextSwitchTest()
}
