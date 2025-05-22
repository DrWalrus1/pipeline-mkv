package eventhandlers

import (
	"io"
	"servermakemkv/makemkv"
	"servermakemkv/makemkv/commands/outputs"
	"servermakemkv/makemkv/streamReader"
)

func MakeMkvInfoEventHandler(reader io.Reader) (standardEventsChannel chan outputs.MakeMkvOutput, discInfoEventChannel chan makemkv.DiscInfo, disconnectChannel chan bool) {
	standardEventsChannel = make(chan outputs.MakeMkvOutput)
	discInfoEventChannel = make(chan makemkv.DiscInfo)
	disconnectChannel = make(chan bool)

	go func() {
		c := streamReader.ParseStream(reader)
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
