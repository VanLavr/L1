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

	go func() {
		data.Write("1", "1")
		data.Write("2", "1")
		data.Write("3", "1")
		data.Write("4", "1")
		data.Write("5", "1")
		data.Write("6", "1")
		data.Write("7", "1")
	}()

	time.Sleep(time.Second)

	for k := range data.data {
		val, ok := data.Read(k)
		if !ok {
			fmt.Println("nothing there")
		}
		fmt.Println(k, val)
	}

	// 2:
	fmt.Println()
	m := make(map[string]string)
	var wg sync.WaitGroup

	for i := 0; i < rand.Intn(10); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			WriteToMap(m, strconv.Itoa(i), strconv.Itoa(rand.Int()))
		}()
	}

	wg.Wait()

	for k := range m {
		val, ok := m[k]
		if !ok {
			fmt.Println("nothing here")
		}
		fmt.Println(k, val)
	}
}

func ReadFromMap(m map[string]string, key string) (string, bool) {
	var mut sync.RWMutex
	// to prevent race condition
	mut.RLock()
	defer mut.RUnlock()

	val, found := m[key]
	return val, found
}

func WriteToMap(m map[string]string, key, value string) {
	var mut sync.RWMutex
	// to prevent race condition
	mut.Lock()
	defer mut.Unlock()

	m[key] = value
}
