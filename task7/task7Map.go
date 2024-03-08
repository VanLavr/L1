package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// first solution:
type MyMap struct {
	data map[string]string
	sync.RWMutex
}

func (m *MyMap) Write(key, value string) {
	// to prevent race condition
	m.Lock()
	defer m.Unlock()

	m.data[key] = value
}

func (m *MyMap) Read(key string) (string, bool) {
	// to prevent race condition
	m.RLock()
	defer m.RUnlock()

	val, found := m.data[key]
	return val, found
}

func main() {
	// 1:
	data := new(MyMap)
	data.data = make(map[string]string)

	for i := 0; i < rand.Intn(100); i++ {
		go func() {
			data.Write(strconv.Itoa(i), strconv.Itoa(rand.Int()))
		}()
	}

	time.Sleep(time.Second)

	for k := range data.data {
		val, ok := data.Read(k)
		if !ok {
			fmt.Println("nothing there")
		}
		fmt.Println(k, val)
	}

	time.Sleep(time.Second * 3)

	// second solution
	lock := sync.RWMutex{}

	b := make(map[string]int)
	b["0"] = 0

	go func(i int) {
		lock.RLock()

		fmt.Println("reading in goroutine", i)
		fmt.Printf("RLock: from goroutine %d: b = %d\n", i, b["0"])
		time.Sleep(time.Second * 3)

		fmt.Printf("RLock: from goroutine %d: lock released\n", i)
		lock.RUnlock()
	}(1)

	go func(i int) {
		lock.Lock()

		b["2"] = i
		fmt.Println("writing in goroutine", i)
		fmt.Printf("Lock: from goroutine %d: b = %d\n", i, b["2"])
		time.Sleep(time.Second * 3)

		fmt.Printf("Lock: from goroutine %d: lock released\n", i)
		lock.Unlock()
	}(2)

	<-time.After(time.Second * 8)

	fmt.Println("*************************************")

	go func(i int) {
		lock.Lock()

		b["3"] = i
		fmt.Println("writing in goroutine", i)
		fmt.Printf("Lock: from goroutine %d: b = %d\n", i, b["3"])
		time.Sleep(time.Second * 3)

		fmt.Printf("Lock: from goroutine %d: lock released\n", i)
		lock.Unlock()
	}(3)

	go func(i int) {
		lock.RLock()

		fmt.Println("reading in goroutine", i)
		fmt.Printf("RLock: from goroutine %d: b = %d\n", i, b["3"])
		time.Sleep(time.Second * 3)

		fmt.Printf("RLock: from goroutine %d: lock released\n", i)
		lock.RUnlock()
	}(4)

	<-time.After(time.Second * 8)
}
