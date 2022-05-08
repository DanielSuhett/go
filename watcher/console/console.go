package console

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func statusOFF() {
	printERR := color.New(color.Bold, color.FgBlack, color.BgRed).PrintfFunc()
	printERR(" OFF ")
}

func statusON() {
	printERR := color.New(color.FgBlack, color.BgGreen).PrintfFunc()
	printERR(" ON ")
}

func Success(url string) {
	fmt.Print(time.Now().Format("02-01-2006 - 15:04:05:	"))
	fmt.Print(url, " ")
	statusON()
	println("")
}

func Failure(url string, status int) {
	fmt.Print(time.Now().Format("02-01-2006 - 15:04:05:	"))
	fmt.Print(url, " ")
	statusOFF()
	println("")
}
