package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/simonswine/go-junit-report/parser"
)

const testStart = "=== RUN   "

var (
	regexStatus = regexp.MustCompile(`--- (PASS|FAIL|SKIP): (.+) \((\d+\.\d+)(?: seconds|s)\)`)
)

func main() {
	out := os.Stdout

	testCh := make(chan *parser.Test)

	go func() {
		_, err := parser.ParseChannel(os.Stdin, "", testCh)
		if err != nil {
			log.Fatal(err)
		}
		close(testCh)
	}()

	for test := range testCh {
		status := ""
		switch test.Result {
		case parser.PASS:
			status = "PASS"
		case parser.SKIP:
			status = "SKIP"
		case parser.FAIL:
			status = "FAIL"
		default:
			status = "UNKNOWN"
		}

		fmt.Fprintf(out, "::group::%s\n", test.Name)
		fmt.Fprintf(out, "%s%s\n", testStart, test.Name)
		fmt.Fprintln(out, strings.Join(test.Output, "\n"))
		fmt.Fprintf(out, "--- %s: %s %s\n", status, test.Name, "todo-duration")
		fmt.Fprintln(out, "::endgroup::", test.Name)
	}
}
