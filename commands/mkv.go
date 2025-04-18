package command

import (
	"log"
	"os/exec"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

// Mkv calls the MakeMKV executable with the given arguments.
func Mkv() {
	cmd := exec.Command("makemkvcon -r --cache=1 info disc:9999")

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
