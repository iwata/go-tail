package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	lineNum, filePath, err := parseArgs()
	if err != nil {
		log.Fatalf("Failed to parse args: %s", err)
	}

	if lineNum < 1 {
		return
	}

	r := getResolver(filePath)
	reader, closer, err := r.resolve()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	o := outPutter{lineNum, reader}
	o.print()
}

func parseArgs() (int, string, error) {
	var n int
	flag.IntVar(&n, "n", 10, "The location is number lines")
	flag.Parse()

	if flag.NArg() != 1 {
		return 0, "", fmt.Errorf("Not support %d args, only one arg", flag.NArg())
	}

	f := flag.Arg(0)
	return n, f, nil
}

func getResolver(filePath string) readResolver {
	return fileReadResolver{filePath}
}

type readResolver interface {
	resolve() (io.Reader, func(), error)
}

type fileReadResolver struct {
	filePath string
}

func (f fileReadResolver) resolve() (io.Reader, func(), error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to open file: %s", err)
	}
	return file, func() { file.Close() }, nil
}

type outPutter struct {
	n int
	r io.Reader
}

func (o outPutter) print() {
	buf := make([]string, 0, o.n)
	scanner := bufio.NewScanner(o.r)
	for scanner.Scan() {
		l := scanner.Text()
		if len(buf) == o.n {
			buf = buf[1:]
		}
		buf = append(buf, l)
	}

	for _, t := range buf {
		fmt.Println(t)
	}
}
