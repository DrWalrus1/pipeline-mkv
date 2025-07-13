package makemkv

import (
	"bufio"
	"io"
	"log"
	"time"
)

func runInitialDiscLoadOnStartup() {
	//TODO: Set this value in config
	initalLoadReader, _, _ := TriggerInitialInfoLoad(time.Minute * 2)
	stringChan := readStream(initalLoadReader)
	go func() {
		for {
			select {
			case newRead, ok := <-stringChan:
				if !ok {
					return
				}
				log.Println(newRead)
			}
		}
	}()
}

func readStream(reader io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
