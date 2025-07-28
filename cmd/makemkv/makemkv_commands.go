package makemkv

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/DrWalrus1/gomakemkv"
)

type IMakeMkvCommandHandler interface {
	LoadConfig(configPath string) error
	TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error)
	TriggerInitialInfoLoad(timeout time.Duration) (io.Reader, context.CancelFunc, error)
	TriggerDiskBackup(decrypt bool, source string, destination string) (io.Reader, context.CancelFunc, error)
	TriggerSaveMkv(source string, title string, destination string) (io.Reader, context.CancelFunc, error)
	RegisterMakeMkv(key string) error
}

type MakeMkvCommandHandler struct {
	ExecutablePath string
}

func (m *MakeMkvCommandHandler) LoadConfig(configPath string) error {
	panic("unimplemented: MakeMkvCommandHandler.LoadConfig")
}

func (m MakeMkvCommandHandler) TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error) {
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

func (m MakeMkvCommandHandler) TriggerInitialInfoLoad(timeoutLength time.Duration) (io.Reader, context.CancelFunc, error) {
	timeoutErr := errors.New("failed to perform initial disc load - timeout")
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
				log.Printf("%s", context.Cause(ctx).Error())

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

func (m MakeMkvCommandHandler) TriggerDiskBackup(decrypt bool, source string, destination string) (io.Reader, context.CancelFunc, error) {
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

func validateSource(source string) error {
	if source == "" {
		return fmt.Errorf("source cannot be empty")
	}

	if strings.HasPrefix(source, "disc:") || strings.HasPrefix(source, "iso:") || strings.HasPrefix(source, "file:") || strings.HasPrefix(source, "dev:") {
		return nil
	}
	return fmt.Errorf("invalid source")
}

func (m MakeMkvCommandHandler) TriggerSaveMkv(source string, title string, destination string) (io.Reader, context.CancelFunc, error) {
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

func (m MakeMkvCommandHandler) RegisterMakeMkv(key string) error {
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
	return gomakemkv.HandleRegisterMakeMkvEvents(outputPipe)

}
