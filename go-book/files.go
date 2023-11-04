package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Dup3() {
	startTime := time.Now()

	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("dup3 took %s\n", elapsedTime)
}

func Dup2() {
	startTime := time.Now()
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("dup2 took %s\n", elapsedTime)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	if input.Err() != nil {
		panic(input.Err().Error())
	}

	for input.Scan() {
		counts[input.Text()]++
	}
}
