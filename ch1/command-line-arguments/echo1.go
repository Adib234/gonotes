// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs := time.Since(start)
	fmt.Println("%.2fs", secs)
	start1 := time.Now()
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)
	secs1 := time.Since(start1)
	fmt.Println("%.2fs", secs1)
}
