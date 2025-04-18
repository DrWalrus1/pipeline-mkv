package stream_test

import (
	"encoding/json"
	"fmt"
	"io"
	"servermakemkv/outputs"
	"servermakemkv/stream"
	"testing"
)

func simulateMakeMkvProgressOutput(t *testing.T) io.Reader {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		maxPercentage := 20
		totalPercentage := 10
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
	c := make(chan outputs.MakeMkvOutput)
	go stream.ParseStream(mockOutput, c)
	for i := range c {
		str, _ := json.Marshal(i)
		t.Log(string(str))
	}
}
