type DiscInfo = {
  name: string
  language: string
  type: string
  titles: TitleInfo[]
}


type TitleInfo = {
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

type VideoInfo = {

}

type AudioInfo = {}
type SubtitleInfo = {}
