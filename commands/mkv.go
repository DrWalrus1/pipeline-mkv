package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"servermakemkv/commands/eventhandlers"
	"servermakemkv/config"
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv"
	"servermakemkv/stream"
	"strings"
)

// MkvInfo calls the MakeMKV executable with the given arguments.
func GetInfo(config *config.Config, source string, stringified chan []byte) {
	cmd := exec.Command("makemkvcon", "-r", "--progress=-stdout", "info", source)

	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error executing makemkvcon: %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
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

	if err := cmd.Wait(); err != nil {
		log.Fatalf("error waiting for command to finish: %s", err.Error())
	}
}

func SaveMkv(source string, title string, destination string, stringified chan []byte) {
	if err := validateSource(source); err != nil {
		stringified <- []byte("Failed to validate")
		close(stringified)
		return
	}

	cmd := exec.Command("makemkvcon", "-r", "--progress=-stdout", "--debug=-stdout", "mkv", source, title, destination)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error executing makemkvcon: %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	events := make(chan outputs.MakeMkvOutput)
	go stream.ParseStream(outputPipe, events)
	for {
		if event, ok := <-events; ok {
			newJson, _ := json.Marshal(event)
			stringified <- newJson
		} else {
			close(stringified)
			break
		}
	}
	if err := cmd.Wait(); err != nil {
		log.Fatalf("error waiting for command to finish: %s", err.Error())
	}
}

func validateSource(source string) error {
	if source == "" {
		return fmt.Errorf("source cannot be empty")
	}

	if strings.HasPrefix(source, "disc:") || strings.HasPrefix(source, "iso:") || strings.HasPrefix(source, "file:") || strings.HasPrefix(source, "dev:") {
		return nil
	}
	return fmt.Errorf("invalid source")
}

func BackupDisk(decrypt bool, source string, destination string, stringified chan []byte, cancelChannel chan bool) {
	var cmd *exec.Cmd
	ctx, cancel := context.WithCancel(context.Background())
	if decrypt {
		cmd = exec.CommandContext(ctx, "makemkvcon", "-r", "backup", "--decrypt", source, destination)
	} else {
		cmd = exec.CommandContext(ctx, "makemkvcon", "-r", "backup", source, destination)
	}
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("error creating stdout pipe: %s", err.Error())
		close(stringified)
		return
	}
	if err := cmd.Start(); err != nil {
		log.Printf("error starting command: %s", err.Error())
		close(stringified)
		return
	}
	go func() {
		<-cancelChannel
		cancel()
	}()
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

	err = cmd.Wait()
	if err != nil {
		// Check if the process was interrupted
		if ctx.Err() == context.Canceled {
			return
		}
		log.Printf("error waiting for command to finish: %s", err.Error())
	}
}

const registerMkvKeyBadKeyPrefix string = "Key not found or invalid"
const registerMkvKeySavedPrefix string = "Registration key saved"

func RegisterMkvKey(key string) int {
	executable := "makemkvcon"
	arguments := "-r"
	command := "reg"
	cmd := exec.Command(executable, arguments, command, key)
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error creating pipe to command. %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("error executing command. %s", err.Error())
	}
	c := make(chan string)
	go stream.ReadStream(outputPipe, c)
	for s := range c {
		switch {
		case strings.HasPrefix(s, registerMkvKeyBadKeyPrefix):
			fmt.Println(s)
			return 400
		case strings.HasPrefix(s, registerMkvKeySavedPrefix):
			fmt.Println(s)
			return 200
		default:
			fmt.Println(s)
		}
	}
	return 500
}
