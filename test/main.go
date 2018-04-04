package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	flag.Parse()
	data := flag.Args()
	if !terminal.IsTerminal(0) {
		reader := bufio.NewReader(os.Stdin)
		for {
			bts, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}

			print("==== " + bts)
			if err == nil {
				data = append(data, string(bts))
			}
		}
	}
	fmt.Println(strings.Join(data, " "))
}
