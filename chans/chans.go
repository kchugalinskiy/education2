package chans

import (
	"fmt"
	"sync"
	"sync/atomic" // https://habr.com/ru/articles/183834/
	"time"
)

// задача - реализовать MyWaitgroup (на каналах)
/*
type MyWaitgroup struct {
	// где-то тут каналы
}

// задача без * - вот этот
func NewMyWaitgroup(delta int) *MyWaitgroup {

}

// задача со * - реализовать этот метод
func (m *MyWaitgroup) Add(delta int) {

}

func (m *MyWaitgroup) Done() {

}

func (m *MyWaitgroup) Wait() {

}
*/

// пример использования
/*
	mwg := NewMyWaitgroup(2)
	mwg.Done()
	mwg.Done()
	mwg.Wait()

	var mwg2 MyWaitgroup
	mwg2.Add(2)
	mwg2.Done()
	mwg2.Add(1)
	mwg2.Done()
	mwg2.Add(2)
	mwg2.Done()
	mwg2.Done()
	mwg2.Done()
	mwg2.Wait()
*/

// задача с ** - генерация массива сейчас работает по времени, переделать на работу до нажатия пользователем комбинации клавиш Ctrl+C
/*
https://go.dev/play/p/tOjXzGQH-s4
https://pkg.go.dev/os/signal
package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
*/

type mutexer struct {
	m    sync.Mutex // https://linux.die.net/man/3/pthread_mutex_unlock
	wg   sync.WaitGroup
	arr  []int
	stop atomic.Bool
}

func NewMutexer() *mutexer {
	m := &mutexer{
		arr: make([]int, 0),
	}
	m.wg.Add(2)
	go m.reader()
	go m.writer()
	return m
}

func (m *mutexer) Close() {
	m.stop.Store(true)
	m.wg.Wait()

	m.arr = nil
}

func (m *mutexer) reader() {
	defer m.wg.Done()
	var maxval int

	for m.stop.Load() == false {
		func() {
			m.m.Lock()
			defer m.m.Unlock()
			for i := range m.arr {
				maxval = i
			}
			m.arr = m.arr[:0]
		}()
	}
	fmt.Printf("mutexer: %d\n", maxval)
}

func (m *mutexer) writer() {
	defer m.wg.Done()
	i := 0
	for m.stop.Load() == false {
		func() {
			m.m.Lock()
			defer m.m.Unlock()
			m.arr = append(m.arr, i)
		}()
		i++
	}
}

type chaneller struct {
	ch     chan int
	exitCh chan struct{}
	exitWg sync.WaitGroup
}

func NewSomeMyStruct() *chaneller {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	c := &chaneller{
		ch:     ch,
		exitWg: wg,
		exitCh: make(chan struct{}),
	}
	go c.reader()
	go c.writer()
	return c
}

func (c *chaneller) writer() {
	defer c.exitWg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
		}
	}()
	defer func() {
		fmt.Println("deferred")
	}()
	i := 0
	for ; ; i++ {
		select {
		case <-c.exitCh:
			fmt.Println("ended")
			return
		default:
			c.ch <- i
		}
	}
}

func (c *chaneller) reader() {
	defer c.exitWg.Done()
	var maxval int
	for i := range c.ch {
		maxval = i
	}
	fmt.Printf("chaneller: %d\n", maxval)
}

func (c *chaneller) Close() error {
	close(c.exitCh)
	close(c.ch)
	c.exitWg.Wait()
	return nil
}

var maxgen atomic.Int32

func generator() {
	maxgen.Add(1)
	fmt.Println("generator", maxgen.Load())
}

func Somefun() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c := NewSomeMyStruct()
		defer c.Close()
		time.Sleep(1 * time.Second)
	}()
	go func() {
		defer wg.Done()
		mtxr := NewMutexer()
		defer mtxr.Close()
		time.Sleep(1 * time.Second)
	}()
	wg.Wait()

	for i := 0; i < 1000; i++ {
		go generator()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println("real max", maxgen.Load())
}
