package makemkv

import "servermakemkv/makemkv/commands/outputs"

type DiscInfo struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	Type     string  `json:"type"`
	Titles   []Title `json:"titles"`
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

type Title struct {
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	Size           string          `json:"size"`
	SizeInBytes    string          `json:"sizeInBytes"`
	Duration       string          `json:"duration"`
	Language       string          `json:"language"`
	Chapters       string          `json:"chapters"`
	OutputFileName string          `json:"outputFileName"`
	VideoTracks    []VideoTrack    `json:"video"`
	AudioTracks    []AudioTrack    `json:"audio"`
	SubtitleTracks []SubtitleTrack `json:"subtitles"`
}

func (t *Title) UpdateTitle(info *outputs.TitleInformation) {
	switch info.AttributeId {
	case outputs.DiskSize:
		t.Size = info.Value
	}

}

type VideoTrack struct {
	Type           string `json:"type"`
	Framerate      string `json:"framerate"`
	VideoSize      string `json:"videoSize"`
	Codec          string `json:"codec"`
	Language       string `json:"language"`
	ConversionType string `json:"conversionType"`
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

type SubtitleTrack struct {
	Type           string `json:"type"`
	Language       string `json:"language"`
	Codec          string `json:"codec"`
	ConversionType string `json:"conversionType"`
}
