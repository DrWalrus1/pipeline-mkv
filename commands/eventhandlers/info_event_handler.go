package eventhandlers

import (
	"io"
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv"
	"servermakemkv/stream"
)

func MakeMkvInfoEventHandler(reader io.Reader, standardEventsChan chan<- outputs.MakeMkvOutput, discInfoEventChan chan<- makemkv.MakeMkvDiscInfo, disconnectChan chan<- bool) {
	c := make(chan outputs.MakeMkvOutput)
	go stream.ParseStream(reader, c)
	var discInfoEvents []outputs.MakeMkvOutput
	for {
		if i, ok := <-c; ok {
			if standardEvent, ok := i.(*outputs.TitleInformation); ok {
				discInfoEvents = append(discInfoEvents, standardEvent)
			} else if standardEvent, ok := i.(*outputs.DiscInformation); ok {
				discInfoEvents = append(discInfoEvents, standardEvent)
			} else if standardEvent, ok := i.(*outputs.StreamInformation); ok {
				discInfoEvents = append(discInfoEvents, standardEvent)
			} else {
				standardEventsChan <- i
			}
		} else {
			if len(discInfoEvents) > 0 {
				discInfoEventChan <- makemkv.MakeMkvOutputsIntoMakeMkvDiscInfo(discInfoEvents)
			}
			break
		}
	}
	disconnectChan <- true
	close(standardEventsChan)
	close(discInfoEventChan)
	close(disconnectChan)
}
