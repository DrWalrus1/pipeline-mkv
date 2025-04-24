package eventhandlers

import (
	"io"
	"servermakemkv/outputs"
	"servermakemkv/stream"
)

func MakeMkvMkvEventHandler(reader io.Reader, events chan outputs.MakeMkvOutput) {
	go stream.ParseStream(reader, events)
}
