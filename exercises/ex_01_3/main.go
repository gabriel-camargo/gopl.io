package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// !+
func main() {
	startEcho3 := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now())

	fmt.Printf("echo3: %.2fs\n", time.Since(startEcho3).Seconds())

	startEcho2 := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("echo2: %.2fs\n", time.Since(startEcho2).Seconds())

}

//!-
