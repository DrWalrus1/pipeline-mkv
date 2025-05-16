package eventhandlers

import (
	"io"
	"servermakemkv/makemkv/commands/outputs"
	"servermakemkv/makemkv/streamReader"
)

func MakeMkvMkvEventHandler(reader io.Reader) <-chan outputs.MakeMkvOutput {
	return streamReader.ParseStream(reader)
}
