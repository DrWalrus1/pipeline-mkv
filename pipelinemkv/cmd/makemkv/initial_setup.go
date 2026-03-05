package makemkv

import (
	"bufio"
	"io"
	"log"
	"time"
)

func runInitialDiscLoadOnStartup(handler IMakeMkvCommandHandler, timeout time.Duration) {
	initalLoadReader, _, _ := handler.TriggerInitialInfoLoad(timeout)
	stringChan := readStream(initalLoadReader)
	go func() {
		for newRead := range stringChan {
			log.Println(newRead)
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
