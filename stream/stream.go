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

func ReadStream(reader io.Reader, c chan<- string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		c <- scanner.Text()
	}
	close(c)
}
