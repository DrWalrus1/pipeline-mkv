package stream

import (
	"bufio"
	"fmt"
	"io"
	"servermakemkv/outputs"
	"servermakemkv/parser"
)

func ParseStream(reader io.Reader, c chan outputs.MakeMkvOutput) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		output, err := parser.Parse(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		}
		c <- output
	}
	close(c)
}
