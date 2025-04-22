package makemkv

import (
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv/ids"
)

type MakeMkvDiscInfo struct {
	Properties map[string]MakeMkvValue `json:"properties"`
	Titles     map[int]MakeMkvTitle    `json:"titles"`
}

type MakeMkvTitle struct {
	Properties map[string]MakeMkvValue         `json:"properties"`
	Streams    map[int]map[string]MakeMkvValue `json:"streams"`
}

type MakeMkvValue struct {
	MessageCodeValue string `json:"messageCodeValue,omitempty"`
	Value            string `json:"value,omitempty"`
}

func MakeMkvOutputsIntoMakeMkvDiscInfo(makemkvOutputs []outputs.MakeMkvOutput) MakeMkvDiscInfo {
	// just grab all the
	mkvDiscInfo := MakeMkvDiscInfo{
		Properties: make(map[string]MakeMkvValue),
		Titles:     MakeMkvOutputsIntoMakeMkvTitles(makemkvOutputs),
	}
	for _, x := range makemkvOutputs {
		if i, ok := x.(*outputs.DiscInformation); ok {
			desc, _ := ids.GetItemAttributeDescription(i.ID)
			messagecode, err := ids.GetAppConstantDescription(i.MessageCodeId)
			if err != nil {
				messagecode = ""
			}
			mkvValue := MakeMkvValue{
				MessageCodeValue: messagecode,
				Value:            i.Value,
			}
			mkvDiscInfo.Properties[desc] = mkvValue
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
					Properties: make(map[string]MakeMkvValue),
					Streams:    make(map[int]map[string]MakeMkvValue),
				}
			}
			desc, _ := ids.GetItemAttributeDescription(i.AttributeId)
			messagecode, err := ids.GetAppConstantDescription(i.MessageCodeId)
			if err != nil {
				messagecode = ""
			}
			mkvValue := MakeMkvValue{
				MessageCodeValue: messagecode,
				Value:            i.Value,
			}
			newTitles[i.TitleIndex].Properties[desc] = mkvValue
		} else if i, ok := x.(*outputs.StreamInformation); ok {
			// if title doesn't exist
			if i.Value == "" {
				continue
			}
			if _, ok := newTitles[i.TitleIndex]; !ok {
				newTitles[i.TitleIndex] = MakeMkvTitle{
					Properties: make(map[string]MakeMkvValue),
					Streams:    make(map[int]map[string]MakeMkvValue),
				}
			}
			// if there are streams but not this one in particular
			if _, ok := newTitles[i.TitleIndex].Streams[i.StreamIndex]; !ok {
				newTitles[i.TitleIndex].Streams[i.StreamIndex] = make(map[string]MakeMkvValue)
			}
			desc, _ := ids.GetItemAttributeDescription(i.AttributeId)
			messagecode, err := ids.GetAppConstantDescription(i.MessageCodeId)
			if err != nil {
				messagecode = ""
			}
			mkvValue := MakeMkvValue{
				MessageCodeValue: messagecode,
				Value:            i.Value,
			}
			newTitles[i.TitleIndex].Streams[i.StreamIndex][desc] = mkvValue
		}

	}
	return newTitles
}
