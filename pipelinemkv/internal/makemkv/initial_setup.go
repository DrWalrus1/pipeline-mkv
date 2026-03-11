package makemkv

import (
	"bufio"
	"log"
	"time"
)

// TODO: should probably turn this into a handler or main function
func runInitialDiscLoadOnStartup(handler IMakeMkvCommandHandler, timeout time.Duration) {
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
