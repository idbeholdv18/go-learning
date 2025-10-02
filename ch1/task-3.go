package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo_v1() {
	var s, sep string

	for i := 0; i < len(os.Args[0:]); i++ {
		s += os.Args[i] + sep
		sep = " "
	}

	fmt.Println(s)
}

func echo_v2() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	start := time.Now()
	echo_v1()
	fmt.Println("echo v1: ", time.Since(start).Microseconds())
	start = time.Now()
	echo_v2()
	fmt.Println("echo v2: ", time.Since(start).Microseconds())
}
