package eventhandlers

import (
	"io"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

func MakeMkvMkvEventHandler(reader io.Reader) <-chan outputs.MakeMkvOutput {
	return stream.ParseStream(reader)
}
