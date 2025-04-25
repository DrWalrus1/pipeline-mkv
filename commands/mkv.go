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
	"strings"
)

type CancellableCommand interface {
	Start() error
	GetStdoutPipe() (io.Reader, error)
	Wait()
	Cancel()
}

type Command struct {
	cmd        *exec.Cmd
	reader     io.Reader
	cancelFunc context.CancelFunc
}

func NewCommand(name string, args ...string) Command {
	ctx, cancel := context.WithCancel(context.Background())
	return Command{
		cmd:        exec.CommandContext(ctx, name, args...),
		cancelFunc: cancel,
	}
}

func (command *Command) Start() error {
	return command.cmd.Start()
}

func (command *Command) GetStdoutPipe() (io.Reader, error) {
	if command.reader == nil {
		output, err := command.cmd.StdoutPipe()
		if err == nil {
			command.reader = output
		}
		return output, err
	}
	return command.reader, nil
}

func (command *Command) Wait() error {
	return command.cmd.Wait()
}

func (command *Command) Cancel() {
	command.cancelFunc()
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
