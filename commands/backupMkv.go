package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

func TriggerDiskBackup(decrypt bool, source string, destination string) (io.Reader, context.CancelFunc, error) {
	var cmd *exec.Cmd
	ctx, cancel := context.WithCancel(context.Background())
	if decrypt {
		cmd = exec.CommandContext(ctx, "makemkvcon", "-r", "backup", "--decrypt", source, destination)
	} else {
		cmd = exec.CommandContext(ctx, "makemkvcon", "-r", "backup", source, destination)
	}
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("error creating stdout pipe: %w", err)
		return nil, cancel, err
	}
	if err := cmd.Start(); err != nil {
		err = fmt.Errorf("error starting command: %w", err)
		return nil, cancel, err
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

func WatchBackupLogs(outputPipe io.Reader, stringified chan []byte) {
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
