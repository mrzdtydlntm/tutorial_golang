package main

import (
	"fmt"
	"runtime"
	"sync"
)

//dengan mutex

type counter struct {
	sync.Mutex
	val int
}

func (c *counter) Add(x int) {
	c.Lock()
	c.val++
	c.Unlock()
}

//tanpa mutex
/*
type counter struct {
	val int
}

func (c *counter) Add(x int) {
	c.val++
}
*/

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var meter counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(meter.Value())
}
