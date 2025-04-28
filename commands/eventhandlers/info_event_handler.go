package eventhandlers

import (
	"io"
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv"
	"servermakemkv/stream"
)

func MakeMkvInfoEventHandler(reader io.Reader) (standardEventsChannel chan outputs.MakeMkvOutput, discInfoEventChannel chan makemkv.MakeMkvDiscInfo, disconnectChannel chan bool) {
	standardEventsChannel = make(chan outputs.MakeMkvOutput)
	discInfoEventChannel = make(chan makemkv.MakeMkvDiscInfo)
	disconnectChannel = make(chan bool)

	go func() {
		c := stream.ParseStream(reader)
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
					standardEventsChannel <- i
				}
			} else {
				if len(discInfoEvents) > 0 {
					discInfoEventChannel <- makemkv.MakeMkvOutputsIntoMakeMkvDiscInfo(discInfoEvents)
				}
				break
			}
		}
		disconnectChannel <- true
		close(standardEventsChannel)
		close(discInfoEventChannel)
		close(disconnectChannel)
	}()
	return standardEventsChannel, discInfoEventChannel, disconnectChannel
}
