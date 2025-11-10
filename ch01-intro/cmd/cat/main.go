package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		copyStream(os.Stdin, os.Stdout)
		return
	}
	for _, fname := range os.Args[1:] {
		if err := catFile(fname); err != nil {
			fmt.Fprintf(os.Stderr, "cat: %s: %v\n", fname, err)
		}
	}
}

func catFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	copyStream(f, os.Stdout)
	return nil
}

func copyStream(r io.Reader, w io.Writer) {
	in := bufio.NewScanner(r)
	for in.Scan() {
		fmt.Fprintln(w, in.Text())
	}
}
