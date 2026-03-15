package makemkv

import (
	"bufio"
	"context"
	"io"
	"log"
	"time"
)

type intialInfoLoadHandler interface {
	TriggerInitialInfoLoad(time.Duration) (io.Reader, context.CancelFunc, error)
}

// TODO: should probably turn this into a handler or main function
func runInitialDiscLoadOnStartup(handler intialInfoLoadHandler, timeout time.Duration) {
	initalLoadReader, _, _ := handler.TriggerInitialInfoLoad(timeout)
	stringChan := make(chan string)

	go func() {
		defer close(stringChan)
		scanner := bufio.NewScanner(initalLoadReader)
		for scanner.Scan() {
			stringChan <- scanner.Text()
		}
	}()

	go func() {
		for newRead := range stringChan {
			log.Println(newRead)
		}
	}()
}
