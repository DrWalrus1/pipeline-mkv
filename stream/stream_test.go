package stream_test

import (
	"fmt"
	"io"
	"servermakemkv/stream"
	"testing"
)

func simulateMakeMkvProgressOutput(t *testing.T) io.Reader {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		maxPercentage := 200
		totalPercentage := 100
		for currentPercentage := 0; currentPercentage <= totalPercentage; currentPercentage++ {
			line := fmt.Sprintf("PRGV:%d,%d,%d\n", currentPercentage, totalPercentage, maxPercentage)
			_, err := writer.Write([]byte(line))
			if err != nil {
				t.Errorf("Error writing to pipe: %v\n", err)
				return
			}
			// Use this if you want to slow down the stream
			// time.Sleep(1 * time.Second)
		}
	}()
	return reader
}

func TestProcessStream(t *testing.T) {
	mockOutput := simulateMakeMkvProgressOutput(t)
	lineHandler := func(line string) {
		t.Log(line)
	}
	err := stream.ProcessStream(mockOutput, lineHandler)
	if err != nil {
		t.Fatalf("processStream returned an error: %v", err)
	}
}
