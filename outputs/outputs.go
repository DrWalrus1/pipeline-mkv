package outputs

type MessageOutput struct {
	Code           string
	Flags          string
	ParameterCount int
	RawMessage     string
	FormatMessage  string
	MessageParams  []string
}

type CurrentProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

type TotalProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

type ProgressBarOutput struct {
	CurrentProgress string
	TotalProgress   string
	MaxProgress     string
}

type DriveScanMessage struct {
	DriveIndex string
	Visible    bool
	Enabled    bool
	Flags      string
	DriveName  string
	DiscName   string
}

type DiscInformationOutputMessage struct {
	TitleCount int
}

type DiscInformation struct {
	ID    string
	Code  string
	Value string
}

type TitleInformation struct {
	ID    string
	Code  string
	Value string
}

type StreamInformation struct {
	ID    string
	Code  string
	Value string
}
