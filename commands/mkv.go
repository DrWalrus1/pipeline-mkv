package command

import (
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

}

func BackupDisk() {

}

func RegisterMkvKey() {

}
