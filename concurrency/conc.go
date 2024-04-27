package concurrency

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var counter int

func printCount() {
	counter++ // 1. чтение counter 2. увеличение counter 3. запись counter
	fmt.Printf("%d\n", counter)
}

func Conc() {
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			printCount()
		}()
	}
}
