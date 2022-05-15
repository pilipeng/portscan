package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, result chan int) {
	for p := range ports {
		// address := fmt.Sprintf("%s:%d", "20.194.168.28", p)
		address := fmt.Sprintf("%s:%d", "123.56.132.6", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}
func main() {
	ports := make(chan int, 100)
	result := make(chan int)
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	go func() {
		for p := 0; p < 100; p++ {
			ports <- p
		}
	}()

	var openports []int
	var closeports []int

	for i := 0; i < 100; i++ {
		p := <-result
		if p == 0 {
			closeports = append(closeports, p)
		} else {
			openports = append(openports, p)
		}
	}
	sort.Ints(openports)
	close(ports)
	close(result)

	for _, p := range closeports {
		fmt.Printf("%d close\n", p)
	}
	for _, p := range openports {
		fmt.Printf("%d open\n", p)
	}
}
