package commands

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os/exec"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

func TriggerSaveMkv(source string, title string, destination string) (io.Reader, context.CancelFunc, error) {
	if err := validateSource(source); err != nil {
		return nil, nil, errors.New("invalid source: " + err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "makemkvcon", "-r", "--progress=-stdout", "--debug=-stdout", "mkv", source, title, destination)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error executing makemkvcon: %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
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

func WatchSaveMkvLogs(outputPipe io.Reader, stringified chan []byte) {
	events := make(chan outputs.MakeMkvOutput)
	go stream.ParseStream(outputPipe, events)
	for {
		event, ok := <-events
		if !ok {
			break
		}
		newJson, err := json.Marshal(event)
		if err != nil {
			log.Printf("error marshaling event: %s", err.Error())
			continue // or break, depending on desired behavior
		}
		stringified <- newJson
	}
	close(stringified) // Ensure stringified is closed after the loop
}