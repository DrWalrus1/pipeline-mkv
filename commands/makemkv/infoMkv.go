package makemkv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"servermakemkv/commands/makemkv/eventhandlers"
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
				newJson, _ := json.Marshal(discEvent)
				stringified <- newJson
			case <-disconnection:
				close(stringified)
				break loop
			}
		}
	}()
	return stringified
}
