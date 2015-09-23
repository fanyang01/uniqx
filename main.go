package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fanyang01/rbtree"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s - Remove duplicated lines and keep the input order for rest of lines\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s [filename]\n", os.Args[0])
	}
	flag.Parse()

	var reader io.Reader
	var err error
	if flag.NArg() == 0 {
		reader = os.Stdin
	} else if flag.NArg() == 1 {
		reader, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	t := rbtree.New(rbtree.CompareString)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := t.Insert(line); ok {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading input:", err)
	}
}
