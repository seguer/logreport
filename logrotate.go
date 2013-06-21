package main

import "os"
import "fmt"
import "github.com/ActiveState/tail"
import "regexp"

func main() {
	t, err := tail.TailFile("/tmp/test/log.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for line := range t.Lines {
		processLine(line)
	}
}

func processLine(line *tail.Line) {
	fmt.Printf("[%s] %s\n", line.Time, line.Text)
}