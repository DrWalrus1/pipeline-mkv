export type DiscInfo = {
  name: string
  language: string
  type: string
  titles: TitleInfo[]
}


export type TitleInfo = {
  id: string
  name: string
  size: string
  sizeInBytes: string
  duration: string
  language: string
  chapters: string
  outputFileName: string
  video: VideoInfo[]
  audio: AudioInfo[]
  subtitles: SubtitleInfo[]
}

export type VideoInfo = {
  type: string
  framework: string
  videoSize: string
  codec: string
  language: string
  conversionType: string
}

export type AudioInfo = {
  type: string
  name: string
  language: string
  bitrate: string
  sampleRate: string
  sampleSize: string
  channelNumbers: string
  conversionType: string
}

export type SubtitleInfo = {
  type: string
  language: string
  codec: string
  conversionType: string
}
