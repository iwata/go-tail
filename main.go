package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var lineNum int
var filePath string

func init() {
	flag.IntVar(&lineNum, "n", 10, "The location is number lines")
}

func main() {
	err := resolveArgs()
	if err != nil {
		log.Fatalf("Failed to resolve args: %s", err)
	}

	if lineNum < 1 {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	buf := make([]string, 0, lineNum)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		if len(buf) == lineNum {
			buf = buf[1:]
		}
		buf = append(buf, l)
	}

	for _, t := range buf {
		fmt.Println(t)
	}
}

func resolveArgs() error {
	flag.Parse()

	if flag.NArg() != 1 {
		return fmt.Errorf("Not support %d args", flag.NArg())
	}

	filePath = flag.Arg(0)
	return nil
}
