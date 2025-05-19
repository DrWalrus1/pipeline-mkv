package makemkv

import (
	"fmt"
	"servermakemkv/makemkv/commands/outputs"
	"strconv"
)

type DiscInfo struct {
	Name     string        `json:"name"`
	Language string        `json:"language"`
	Type     string        `json:"type"`
	Titles   map[int]Title `json:"titles"`
}

func NewDisc() DiscInfo {
	return DiscInfo{
		Titles: make(map[int]Title),
	}
}

func (di *DiscInfo) UpdateDiscInfo(info outputs.DiscInformation) {
	switch info.ID {
	case outputs.Name:
		di.Name = info.Value
	case outputs.MetadataLanguageName:
		di.Language = info.Value
	case outputs.Type:
		di.Type = info.Value
	}
}

func (di *DiscInfo) UpsertDiscTitleMetadata(info outputs.TitleInformation) {
	title, exists := di.Titles[info.TitleIndex]
	if !exists {
		title = NewTitle(strconv.Itoa(info.TitleIndex))
	}
	title.UpdateTitle(info)
	di.Titles[info.TitleIndex] = title
}

func (di *DiscInfo) UpsertTitleStreamMetadata(info outputs.StreamInformation) {
	title, exists := di.Titles[info.TitleIndex]
	if !exists {
		title = NewTitle(strconv.Itoa(info.TitleIndex))
	}
	title.UpsertStreamData(info)
	di.Titles[info.TitleIndex] = title
}

type Title struct {
	Id             string                `json:"id"`
	Name           string                `json:"name"`
	Size           string                `json:"size"`
	SizeInBytes    string                `json:"sizeInBytes"`
	Duration       string                `json:"duration"`
	Language       string                `json:"language"`
	Chapters       string                `json:"chapters"`
	OutputFileName string                `json:"outputFileName"`
	VideoTracks    map[int]VideoTrack    `json:"video"`
	AudioTracks    map[int]AudioTrack    `json:"audio"`
	SubtitleTracks map[int]SubtitleTrack `json:"subtitles"`
}

func NewTitle(id string) Title {
	return Title{
		Id:             id,
		VideoTracks:    make(map[int]VideoTrack),
		AudioTracks:    make(map[int]AudioTrack),
		SubtitleTracks: make(map[int]SubtitleTrack),
	}
}

func (t *Title) UpdateTitle(info outputs.TitleInformation) {
	switch info.AttributeId {
	case outputs.Name:
		t.Name = info.Value
	case outputs.DiskSize:
		t.Size = info.Value
	case outputs.DiskSizeBytes:
		t.SizeInBytes = info.Value
	case outputs.Duration:
		t.Duration = info.Value
	case outputs.MetadataLanguageName:
		t.Language = info.Value
	case outputs.ChapterCount:
		t.Chapters = info.Value
	case outputs.OutputFileName:
		t.OutputFileName = info.Value
	}
}

func (t *Title) UpsertStreamData(info outputs.StreamInformation) {
	// Create new stream if the type is detected
	if info.AttributeId == outputs.Type {
		if info.Value == "Video" {
			t.VideoTracks[info.StreamIndex] = VideoTrack{
				Type: info.Value,
			}
		} else if info.Value == "Audio" {
			t.AudioTracks[info.StreamIndex] = AudioTrack{
				Type: info.Value,
			}
		} else {
			t.SubtitleTracks[info.StreamIndex] = SubtitleTrack{
				Type: info.Value,
			}
		}
	}
	videoTrack, ok := t.VideoTracks[info.StreamIndex]
	if ok {
		videoTrack.UpdateVideoTrack(info)
		t.VideoTracks[info.StreamIndex] = videoTrack
		return
	}
	audioTrack, ok := t.AudioTracks[info.StreamIndex]
	if ok {
		audioTrack.UpdateAudioTrack(info)
		t.AudioTracks[info.StreamIndex] = audioTrack
		return
	}

	subtitleTrack, ok := t.SubtitleTracks[info.StreamIndex]
	if ok {
		subtitleTrack.UpdateSubtitleTrack(info)
		t.SubtitleTracks[info.StreamIndex] = subtitleTrack
		return
	}
	panic(fmt.Sprintf("Attempted to parse out of order stream information"))
	// TODO: CONSIDER making the array a queue. if we don't find the type first skip for now and re-enqueue

}

type VideoTrack struct {
	Type           string `json:"type"`
	Framerate      string `json:"framerate"`
	VideoSize      string `json:"videoSize"`
	Codec          string `json:"codec"`
	Language       string `json:"language"`
	ConversionType string `json:"conversionType"`
}

func (vt *VideoTrack) UpdateVideoTrack(info outputs.StreamInformation) {
	switch info.AttributeId {
	case outputs.VideoFrameRate:
		vt.Framerate = info.Value
	case outputs.VideoSize:
		vt.VideoSize = info.Value
	case outputs.CodecShort:
		vt.Codec = info.Value
	case outputs.MetadataLanguageName:
		vt.Language = info.Value
	case outputs.OutputConversionType:
		vt.ConversionType = info.Value
	}
}

type AudioTrack struct {
	Type           string `json:"type"`
	Name           string `json:"name"`
	Language       string `json:"language"`
	Bitrate        string `json:"bitrate"`
	SampleRate     string `json:"sampleRate"`
	SampleSize     string `json:"sampleSize"`
	ChannelNumbers string `json:"channelNumbers"`
	ConversionType string `json:"conversionType"`
}

func (vt *AudioTrack) UpdateAudioTrack(info outputs.StreamInformation) {
	switch info.AttributeId {
	case outputs.Name:
		vt.Name = info.Value
	case outputs.MetadataLanguageName:
		vt.Language = info.Value
	case outputs.Bitrate:
		vt.Bitrate = info.Value
	case outputs.AudioSampleRate:
		vt.SampleRate = info.Value
	case outputs.AudioSampleSize:
		vt.SampleSize = info.Value
	case outputs.AudioChannelsCount:
		vt.ChannelNumbers = info.Value
	case outputs.OutputConversionType:
		vt.ConversionType = info.Value
	}
}

type SubtitleTrack struct {
	Type           string `json:"type"`
	Language       string `json:"language"`
	Codec          string `json:"codec"`
	ConversionType string `json:"conversionType"`
}

func (vt *SubtitleTrack) UpdateSubtitleTrack(info outputs.StreamInformation) {
	switch info.AttributeId {
	case outputs.MetadataLanguageName:
		vt.Language = info.Value
	case outputs.CodecShort:
		vt.Codec = info.Value
	case outputs.OutputConversionType:
		vt.ConversionType = info.Value
	}

}

func MakeMkvOutputsIntoMakeMkvDiscInfo(makemkvOutputs []outputs.MakeMkvOutput) DiscInfo {
	mkvDiscInfo := NewDisc()
	for _, x := range makemkvOutputs {
		if i, ok := x.(*outputs.DiscInformation); ok {
			mkvDiscInfo.UpdateDiscInfo(*i)
		} else if i, ok := x.(*outputs.TitleInformation); ok {
			mkvDiscInfo.UpsertDiscTitleMetadata(*i)
		} else if i, ok := x.(*outputs.StreamInformation); ok {
			mkvDiscInfo.UpsertTitleStreamMetadata(*i)
		}
	}
	return mkvDiscInfo
}
