package streamtracker_test

import (
	"context"
	"io"
	"pipelinemkv/cmd/streamTracker"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var dummyCancelFunc context.CancelFunc = func() {}

func TestStreamTracker(t *testing.T) {
	streamTracker := streamtracker.NewStreamTracker()
	originalReaderInput := io.Reader(strings.NewReader("Hello world"))

	streamTracker.AddStream("test", &originalReaderInput, dummyCancelFunc)

	var output [5]byte

	originalReaderInput.Read(output[:])
	var outputString = string(output[:])

	assert.Equal(t, outputString, "Hello")

	reattachedStream, _ := streamTracker.GetStream("test")
	var reattachedOutput [5]byte
	(*reattachedStream).Read(reattachedOutput[:])
	outputString = string(reattachedOutput[:])
	assert.Equal(t, outputString, " worl")

}
