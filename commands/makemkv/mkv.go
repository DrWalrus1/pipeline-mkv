package makemkv

import (
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"servermakemkv/commands/makemkv/eventhandlers"
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
