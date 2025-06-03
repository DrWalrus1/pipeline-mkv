package commands

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"pipelinemkv/makemkv/commands/eventhandlers"
	"pipelinemkv/makemkv/commands/outputs"
	"time"
)

func TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "makemkvcon", "-r", "--progress=-stdout", "info", source)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error creating stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error starting command: %w", err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.Canceled {
				return
			}
			cancel()
			log.Printf("error waiting for command: %s", err.Error())
		}
	}()
	return outputPipe, cancel, nil
}

// MkvInfo calls the MakeMKV executable with the given arguments.
func WatchInfoLogs(outputPipe io.Reader) <-chan []byte {
	stringified := make(chan []byte)
	go func() {

		standardEvents, discEvents, disconnection := eventhandlers.MakeMkvInfoEventHandler(outputPipe)

	loop:
		for {
			select {
			case standardEvent := <-standardEvents:
				newJson, _ := json.Marshal(standardEvent)
				stringified <- newJson
			case discEvent := <-discEvents:
				newJson, _ := json.Marshal(outputs.JsonWrapper{
					Type: discEvent.GetTypeName(),
					Data: discEvent,
				})
				stringified <- newJson
			case <-disconnection:
				close(stringified)
				break loop
			}
		}
	}()
	return stringified
}

func TriggerInitialInfoLoad(timeoutLength time.Duration) (io.Reader, context.CancelFunc, error) {
	timeoutErr := errors.New("Failed to perform initial disc load - Timeout")
	ctx, cancel := context.WithTimeoutCause(context.Background(), timeoutLength, timeoutErr)

	cmd := exec.CommandContext(ctx, "makemkvcon", "-r", "--cache=1", "--progress=-stdout", "info", "disc:9999")
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error creating stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		cancel()
		return nil, nil, fmt.Errorf("error starting command: %w", err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			if context.Cause(ctx) == timeoutErr {
				log.Printf(context.Cause(ctx).Error())

				return
			}
			if ctx.Err() == context.Canceled {
				return
			}
			cancel()
			log.Printf("error waiting for command: %s\n%s\n", cmd.String(), err.Error())
		}
	}()
	return outputPipe, cancel, nil

}
