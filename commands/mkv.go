package commands

import (
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

func BackupDisk(source string) {
	destination := "./"
	exec.Command("makemkvcon", "-r", "backup", source, destination)
}

/*
	only two outputs

--- SUCCESS ---
Found registration key  : <actual key>
Registration key saved.
---------------
--- FAIL ---
Key not found or invalid
------------
*/

const registerMkvKeySuccessPrefix string = "Found registration key"
const registerMkvKeyKeyAlreadyExistsPrefix string = "Current registration key"
const registerMkvKeyBadKeyPrefix string = "Key not found or invalid"

func RegisterMkvKey(key string) error {
	executable := "makemkvcon"
	arguments := "-r"
	command := "reg"
	cmd := exec.Command(fmt.Sprintf("%s %s %s %s", executable, arguments, command, key))
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating pipe to command. %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error executing command. %w", err)
	}
	c := make(chan string)
	go stream.ReadStream(outputPipe, c)
	for s := range c {
		fmt.Println(s)
	}
	return nil
}
