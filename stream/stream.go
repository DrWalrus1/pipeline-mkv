package stream

import (
	"bufio"
	"fmt"
	"io"
	"servermakemkv/outputs"
	"servermakemkv/parser"
)

func ParseStream(reader io.Reader) <-chan outputs.MakeMkvOutput {
	c := make(chan outputs.MakeMkvOutput)
	go func() {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output, err := parser.Parse(scanner.Text())
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			c <- output
		}
		close(c)
	}()
	return c
}

func ReadStream(reader io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println("hello")
			c <- scanner.Text()
		}
	}()
	return c
}
