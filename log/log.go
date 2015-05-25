package log

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"time"
)

var green = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var white = color.New(color.FgWhite).Add(color.Bold).PrintfFunc()
var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
var italic = color.New().Add(color.Italic).SprintFunc()

func Standard(message string) {
	white(message + "\n")
}

func Success(method string, message string, startTime time.Time) {
	// elapsed time
	elapsed := time.Since(startTime)

	// printf
	fmt.Printf("[%s] %s %s - %s.\n", green(http.StatusOK), green(method), message, italic(elapsed.String()))
}

func Warning(method string, message string, startTime time.Time) {
	// elapsed time
	elapsed := time.Since(startTime)
	// printf
	fmt.Printf("[%s] %s %s - %s.\n", yellow(http.StatusNotFound), yellow(method), message, italic(elapsed.String()))
}

func Danger(method string, message string, startTime time.Time) {
	// elapsed time
	elapsed := time.Since(startTime)
	// printf
	fmt.Printf("[%s] %s %s - %s.\n", red(http.StatusInternalServerError), red(method), message, italic(elapsed.String()))
}
