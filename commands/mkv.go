package command

import (
	"log"
	"os/exec"
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
	stream.ProcessStream(outputPipe, handleLine)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	log.Println("Hello2")
}

func handleLine(line string) {
	log.Println(line)
}
