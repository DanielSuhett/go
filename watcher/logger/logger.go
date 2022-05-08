package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
)

const permissions = os.O_RDWR | os.O_CREATE | os.O_APPEND
const path = "logger/logs/"
const extension = ".log"

func Write(site string, state int) {
	file, err := os.OpenFile(path+site+extension, permissions, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02-01-2006 - 15:04:05:	") + site + " STATUS: " + strconv.Itoa(state) + "\n")

	file.Close()
}

func Picking(sites []string) {
	var command int

	fmt.Println("Choose your log file")
	for index, site := range sites {
		fmt.Println(strconv.Itoa(index) + " - " + site)
	}

	fmt.Scan(&command)

	if command > len(sites) || reflect.TypeOf(command).Kind() != reflect.Int {
		fmt.Println("UNKNOWN OPTION")
		os.Exit(-1)
	}

	Read(sites[command])
}

func Read(site string) {
	file, err := ioutil.ReadFile(path + site + extension)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
