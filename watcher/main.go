package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	console "watch/console"
	logger "watch/logger"

	"github.com/inancgumus/screen"
)

const delay = 5

func main() {
	renderMenu()
	readFileSites()

	command := readCommand()
	screen.Clear()

	switch command {
	case 1:
		for {
			Watch()
		}
	case 2:
		sites := readFileSites()
		logger.Picking(sites)
	case 0:
		fmt.Println("Exit")
		os.Exit(0)
	default:
		os.Exit(-1)
	}
}

func readFileSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func renderMenu() {
	fmt.Println("1 - Watch")
	fmt.Println("2 - Logs")
	fmt.Println("0 - Exit")
}

func Watch() {
	sites := readFileSites()

	for _, site := range sites {
		ping(site)
	}

	println("")
	time.Sleep(delay * time.Second)
}

func ping(site string) {
	res, err := http.Get("https://" + site)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if res.StatusCode == 200 {
		console.Success(site)
	} else {
		console.Failure(site, res.StatusCode)
	}

	logger.Write(site, res.StatusCode)
}
