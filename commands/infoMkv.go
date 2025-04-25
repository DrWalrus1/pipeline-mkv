package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"servermakemkv/commands/eventhandlers"
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv"
)

func TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "makemkvcon", "-r", "--progress=-stdout", "info", source)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, cancel, fmt.Errorf("error creating stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return nil, cancel, fmt.Errorf("error starting command: %w", err)
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.Canceled {
				return
			}
			log.Printf("error waiting for command: %s", err.Error())
		}
	}()
	return outputPipe, cancel, nil
}

// MkvInfo calls the MakeMKV executable with the given arguments.
func WatchInfoLogs(outputPipe io.Reader, stringified chan<- []byte) {
	standardEvents := make(chan outputs.MakeMkvOutput)
	discEvents := make(chan makemkv.MakeMkvDiscInfo)
	disconnection := make(chan bool)

	go eventhandlers.MakeMkvInfoEventHandler(outputPipe, standardEvents, discEvents, disconnection)

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
}
