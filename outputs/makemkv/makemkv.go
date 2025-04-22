package makemkv

import (
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv/ids"
)

type MakeMkvDiscInfo struct {
	Properties map[string]string    `json:"properties"`
	Titles     map[int]MakeMkvTitle `json:"titles"`
}

type MakeMkvTitle struct {
	Properties map[string]string         `json:"properties"`
	Streams    map[int]map[string]string `json:"streams"`
}

func MakeMkvOutputsIntoMakeMkvDiscInfo(makemkvOutputs []outputs.MakeMkvOutput) MakeMkvDiscInfo {
	// just grab all the
	mkvDiscInfo := MakeMkvDiscInfo{
		Properties: make(map[string]string),
		Titles:     MakeMkvOutputsIntoMakeMkvTitles(makemkvOutputs),
	}
	for _, x := range makemkvOutputs {
		if i, ok := x.(*outputs.DiscInformation); ok {
			desc, _ := ids.GetItemAttributeDescription(i.ID)
			mkvDiscInfo.Properties[desc] = i.Value
		}
	}
	return mkvDiscInfo
}

func MakeMkvOutputsIntoMakeMkvTitles(makeMkvOutputs []outputs.MakeMkvOutput) map[int]MakeMkvTitle {
	newTitles := make(map[int]MakeMkvTitle)

	for _, x := range makeMkvOutputs {
		if i, ok := x.(*outputs.TitleInformation); ok {
			// new title detected
			if _, ok := newTitles[i.TitleIndex]; !ok {
				newTitles[i.TitleIndex] = MakeMkvTitle{
					Properties: make(map[string]string),
					Streams:    make(map[int]map[string]string),
				}
			}
			desc, _ := ids.GetItemAttributeDescription(i.AttributeId)
			newTitles[i.TitleIndex].Properties[desc] = i.Value
		} else if i, ok := x.(*outputs.StreamInformation); ok {
			// if title doesn't exist
			if _, ok := newTitles[i.TitleIndex]; !ok {
				newTitles[i.TitleIndex] = MakeMkvTitle{
					Properties: make(map[string]string),
					Streams:    make(map[int]map[string]string),
				}
			}
			// if there are streams but not this one in particular
			if _, ok := newTitles[i.TitleIndex].Streams[i.StreamIndex]; !ok {
				newTitles[i.TitleIndex].Streams[i.StreamIndex] = make(map[string]string)
			}
			desc, _ := ids.GetItemAttributeDescription(i.AttributeId)
			newTitles[i.TitleIndex].Streams[i.StreamIndex][desc] = i.Value
		}

	}
	return newTitles
}
