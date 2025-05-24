package commands

import (
	"fmt"
	"log"
	"os/exec"
	"pipelinemkv/makemkv/commands/eventhandlers"
	"strings"
)

func validateSource(source string) error {
	if source == "" {
		return fmt.Errorf("source cannot be empty")
	}

	if strings.HasPrefix(source, "disc:") || strings.HasPrefix(source, "iso:") || strings.HasPrefix(source, "file:") || strings.HasPrefix(source, "dev:") {
		return nil
	}
	return fmt.Errorf("invalid source")
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
	return eventhandlers.HandleRegisterMakeMkvEvents(outputPipe)
}
