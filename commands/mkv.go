package commands

import (
	"fmt"
	"log"
	"os/exec"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

// MkvInfo calls the MakeMKV executable with the given arguments.
func GetInfo() {
	cmd := exec.Command("makemkvcon -r info disc:9999")

	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("error executing makemkvcon: %s", err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// TODO:
	c := make(chan outputs.MakeMkvOutput)
	go stream.ParseStream(outputPipe, c)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func SaveMkv() {
	exec.Command("makemkvcon -r mkv <source> <title_id> <destination> disc:0")

}

func BackupDisk() {
	exec.Command("makemkvcon -r backup <source> <destination>")

}

func RegisterMkvKey(key string) error {
	executable := "makemkvcon"
	arguments := "-r"
	command := "reg"
	cmd := exec.Command(fmt.Sprintf("%s %s %s %s", executable, arguments, command, key))
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error creating pipe to command. %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error executing command. %w", err)
	}
	c := make(chan string)
	go stream.ReadStream(outputPipe, c)
	for s := range c {
		fmt.Println(s)
	}
	return nil
}
