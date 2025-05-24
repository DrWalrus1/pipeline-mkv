package eventhandlers

import (
	"io"
	"pipelinemkv/makemkv/commands/outputs"
	"pipelinemkv/makemkv/streamReader"
)

func MakeMkvMkvEventHandler(reader io.Reader) <-chan outputs.MakeMkvOutput {
	return streamReader.ParseStream(reader)
}
