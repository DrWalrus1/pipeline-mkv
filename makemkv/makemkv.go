package makemkv

import "pipelinemkv/makemkv/commands/outputs"

type MakeMkvDiscInfo struct {
	Properties map[string]MakeMkvValue `json:"properties"`
	Titles     map[int]MakeMkvTitle    `json:"titles"`
}

func (mkvDiscInfo *MakeMkvDiscInfo) addDiscInfoVerbose(discInfo outputs.DiscInformation) {
	desc, _ := outputs.GetItemAttributeDescription(discInfo.ID)
	messagecode, err := outputs.GetAppConstantDescription(discInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            discInfo.Value,
	}
	mkvDiscInfo.Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addTitleInformationVerbose(titleInfo outputs.TitleInformation) {
	if _, ok := discInfo.Titles[titleInfo.TitleIndex]; !ok {
		discInfo.Titles[titleInfo.TitleIndex] = MakeMkvTitle{
			Properties: make(map[string]MakeMkvValue),
			Streams:    make(map[int]map[string]MakeMkvValue),
		}
	}
	desc, _ := outputs.GetItemAttributeDescription(titleInfo.AttributeId)
	messagecode, err := outputs.GetAppConstantDescription(titleInfo.MessageCodeId)
	if err != nil {
		messagecode = ""
	}
	mkvValue := MakeMkvValue{
		MessageCodeValue: messagecode,
		Value:            titleInfo.Value,
	}
	discInfo.Titles[titleInfo.TitleIndex].Properties[desc] = mkvValue
}

func (discInfo *MakeMkvDiscInfo) addStreamInformationVerbose(streamInfo outputs.StreamInformation) {
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
	desc, _ := outputs.GetItemAttributeDescription(streamInfo.AttributeId)
	messagecode, err := outputs.GetAppConstantDescription(streamInfo.MessageCodeId)
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

func MakeMkvOutputsIntoMakeMkvDiscInfoVerbose(makemkvOutputs []outputs.MakeMkvOutput) MakeMkvDiscInfo {
	mkvDiscInfo := MakeMkvDiscInfo{
		Properties: make(map[string]MakeMkvValue),
		Titles:     make(map[int]MakeMkvTitle),
	}
	for _, x := range makemkvOutputs {
		if i, ok := x.(*outputs.DiscInformation); ok {
			mkvDiscInfo.addDiscInfoVerbose(*i)
		} else if i, ok := x.(*outputs.TitleInformation); ok {
			mkvDiscInfo.addTitleInformationVerbose(*i)
		} else if i, ok := x.(*outputs.StreamInformation); ok {
			mkvDiscInfo.addStreamInformationVerbose(*i)
		}
	}
	return mkvDiscInfo
}
