package outputs

type MakeMkvOutput interface {
	GetTypeName() string
}

type JsonWrapper struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type MessageOutput struct {
	Code           string
	Flags          string
	ParameterCount int
	RawMessage     string
	FormatMessage  string
	MessageParams  []string
}

func (mg MessageOutput) GetTypeName() string {
	return "MessageOutput"
}

type CurrentProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

func (c CurrentProgressTitleOutput) GetTypeName() string {
	return "CurrentProgressTitleOutput"
}

type TotalProgressTitleOutput struct {
	Code string
	ID   string
	Name string
}

func (mg TotalProgressTitleOutput) GetTypeName() string {
	return "TotalProgressTitleOutput"
}

type ProgressBarOutput struct {
	CurrentProgress string
	TotalProgress   string
	MaxProgress     string
}

func (mg ProgressBarOutput) GetTypeName() string {
	return "ProgressBarOutput"
}

type DriveScanMessage struct {
	DriveIndex string
	Visible    bool
	Enabled    bool
	Flags      string
	DriveName  string
	DiscName   string
}

func (mg DriveScanMessage) GetTypeName() string {
	return "DriveScanMessage"
}

type DiscInformationOutputMessage struct {
	TitleCount int
}

func (mg DiscInformationOutputMessage) GetTypeName() string {
	return "DiscInformationOutputMessage"
}

type DiscInformation struct {
	ID    string
	Code  string
	Value string
}

func (mg DiscInformation) GetTypeName() string {
	return "DiscInformation"
}

type TitleInformation struct {
	ID    string
	Code  string
	Value string
}

func (c TitleInformation) GetTypeName() string {
	return "TitleInformation"
}

type StreamInformation struct {
	ID    string
	Code  string
	Value string
}

func (c StreamInformation) GetTypeName() string {
	return "StreamInformation"
}
