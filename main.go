package main

import (
	"fmt"
	"sync"
	"time"

	"./scanner"
)

var portFrom, portTo int

func main() {
	portFrom = 1
	portTo = 1024
	fmt.Println("Sync VS Async")
	checkAsync("udp", "localhost")
	checkSync("udp", "localhost")
}

func checkAsync(protocol, hostname string) {
	start := time.Now()
	var result []int

	var wg sync.WaitGroup
	var m sync.Mutex
	for i := portFrom; i < portTo; i++ {
		wg.Add(1)
		go scanAsync(protocol, hostname, i, &result, &m, &wg)
	}
	wg.Wait()
	fmt.Println("Async took ", time.Since(start), " found open ports: ", len(result))
}

func scanAsync(protocol, hostname string, port int, result *[]int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	err := scanner.Scan(protocol, hostname, port)
	if err == nil {
		m.Lock()
		*result = append(*result, port)
		m.Unlock()
	}
}

func checkSync(protocol, hostname string) {
	start := time.Now()
	var result []int
	for i := portFrom; i < portTo; i++ {
		err := scanner.Scan(protocol, hostname, i)
		if err == nil {
			result = append(result, i)
		}
	}
	fmt.Println("Sync took ", time.Since(start), " found open ports: ", len(result))
}
