package makemkv

import (
	"servermakemkv/outputs"
	"servermakemkv/outputs/makemkv/ids"
)

type MakeMkvDiscInfo struct {
	Properties map[string]MakeMkvValue `json:"properties"`
	Titles     map[int]MakeMkvTitle    `json:"titles"`
}

func (mkvDiscInfo *MakeMkvDiscInfo) addDiscInfo(discInfo outputs.DiscInformation) {
	desc, _ := ids.GetItemAttributeDescription(discInfo.ID)
	messagecode, err := ids.GetAppConstantDescription(discInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            discInfo.Value,
	}
	mkvDiscInfo.Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addTitleInformation(titleInfo outputs.TitleInformation) {
	if _, ok := discInfo.Titles[titleInfo.TitleIndex]; !ok {
		discInfo.Titles[titleInfo.TitleIndex] = MakeMkvTitle{
			Properties: make(map[string]MakeMkvValue),
			Streams:    make(map[int]map[string]MakeMkvValue),
		}
	}
	desc, _ := ids.GetItemAttributeDescription(titleInfo.AttributeId)
	messagecode, err := ids.GetAppConstantDescription(titleInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            titleInfo.Value,
	}
	discInfo.Titles[titleInfo.TitleIndex].Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addStreamInformation(streamInfo outputs.StreamInformation) {
	if streamInfo.Value == "" {
		return
	}
	if _, ok := discInfo.Titles[streamInfo.TitleIndex]; !ok {
		discInfo.Titles[streamInfo.TitleIndex] = MakeMkvTitle{
			Properties: make(map[string]MakeMkvValue),
			Streams:    make(map[int]map[string]MakeMkvValue),
		}
	}
	// if there are streams but not this one in particular
	if _, ok := discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex]; !ok {
		discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex] = make(map[string]MakeMkvValue)
	}
	desc, _ := ids.GetItemAttributeDescription(streamInfo.AttributeId)
	messagecode, err := ids.GetAppConstantDescription(streamInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            streamInfo.Value,
	}
	discInfo.Titles[streamInfo.TitleIndex].Streams[streamInfo.StreamIndex][desc] = mkvValue
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
	mkvDiscInfo := MakeMkvDiscInfo{
		Properties: make(map[string]MakeMkvValue),
		Titles:     make(map[int]MakeMkvTitle),
	}
	for _, x := range makemkvOutputs {
		if i, ok := x.(*outputs.DiscInformation); ok {
			mkvDiscInfo.addDiscInfo(*i)
		} else if i, ok := x.(*outputs.TitleInformation); ok {
			mkvDiscInfo.addTitleInformation(*i)
		} else if i, ok := x.(*outputs.StreamInformation); ok {
			mkvDiscInfo.addStreamInformation(*i)
		}
	}
	return mkvDiscInfo
}
