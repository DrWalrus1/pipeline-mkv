package ids

import "errors"

var attributeDetailedDescription = map[AppItemAttributeId]string{
	Unknown:                      "Unknown",
	Type:                         "Type",
	Name:                         "Name",
	LangCode:                     "LanguageCode",
	LangName:                     "LanguageName",
	CodecId:                      "CodecID",
	CodecShort:                   "ShortCodecName",
	CodecLong:                    "LongCodecName",
	ChapterCount:                 "NumberOfChapters",
	Duration:                     "Duration",
	DiskSize:                     "DiskSize",
	DiskSizeBytes:                "DiskSizeInBytes",
	StreamTypeExtension:          "StreamTypeExtension",
	Bitrate:                      "Bitrate",
	AudioChannelsCount:           "NumberOfAudioChannels",
	AngleInfo:                    "AngleInformation",
	SourceFileName:               "SourceFileName",
	AudioSampleRate:              "AudioSampleRate",
	AudioSampleSize:              "AudioSampleSize",
	VideoSize:                    "VideoSize",
	VideoAspectRatio:             "VideoAspectRatio",
	VideoFrameRate:               "VideoFrameRate",
	StreamFlags:                  "StreamFlags",
	DateTime:                     "DateAndTime",
	OriginalTitleId:              "OriginalTitleID",
	SegmentsCount:                "NumberofSegments",
	SegmentsMap:                  "SegmentsMap",
	OutputFileName:               "OutputFileName",
	MetadataLanguageCode:         "MetadataLanguageCode",
	MetadataLanguageName:         "MetadataLanguageName",
	TreeInfo:                     "TreeInformation",
	PanelTitle:                   "PanelTitle",
	VolumeName:                   "VolumeName",
	OrderWeight:                  "OrderWeight",
	OutputFormat:                 "OutputFormat",
	OutputFormatDescription:      "OutputFormatDescription",
	SeamlessInfo:                 "SeamlessInformation",
	PanelText:                    "PanelText",
	MkvFlags:                     "MKVFlags",
	MkvFlagsText:                 "MKVFlagsText",
	AudioChannelLayoutName:       "AudioChannelLayoutName",
	OutputCodecShort:             "OutputShortCodecName",
	OutputConversionType:         "OutputConversionType",
	OutputAudioSampleRate:        "OutputAudioSampleRate",
	OutputAudioSampleSize:        "OutputAudioSampleSize",
	OutputAudioChannelsCount:     "OutputNumberOfAudioChannels",
	OutputAudioChannelLayoutName: "OutputAudioChannelLayout Name",
	OutputAudioChannelLayout:     "OutputAudioChannelLayout",
	OutputAudioMixDescription:    "OutputAudioMixDescription",
	Comment:                      "Comment",
	OffsetSequenceId:             "OffsetSequenceID",
	MaxValue:                     "MaxValue",
}

func GetItemAttributeDescription(id int) (string, error) {
	if desc, ok := attributeDetailedDescription[id]; ok {
		return desc, nil
	}
	return "", errors.New("Unknown Application Item Attribute")
}

const (
	AP_DskFsFlagDvdFilesPresent    int = 1
	AP_DskFsFlagHdvdFilesPresent   int = 2
	AP_DskFsFlagBlurayFilesPresent int = 4
	AP_DskFsFlagAacsFilesPresent   int = 8
	AP_DskFsFlagBdsvmFilesPresent  int = 16
)

var diskFileFlagsDescriptions = map[int]string{
	AP_DskFsFlagDvdFilesPresent:    "DVD files present on disk",
	AP_DskFsFlagHdvdFilesPresent:   "HD DVD files present on disk",
	AP_DskFsFlagBlurayFilesPresent: "Blu-ray files present on disk",
	AP_DskFsFlagAacsFilesPresent:   "Aacs files present on disk",
	AP_DskFsFlagBdsvmFilesPresent:  "Blu-ray disc movie folder files present on disk",
}

func GetDiskFileFlagDescription(id int) string {
	if desc, ok := diskFileFlagsDescriptions[id]; ok {
		return desc
	}
	return "Unknown"
}

const (
	AP_DriveStateNoDrive     uint = 256
	AP_DriveStateUnmounting  uint = 257
	AP_DriveStateEmptyClosed uint = 0
	AP_DriveStateEmptyOpen   uint = 1
	AP_DriveStateInserted    uint = 2
	AP_DriveStateLoading     uint = 3
)

var driveStateDescriptions = map[uint]string{
	AP_DriveStateNoDrive:     "No Drive detected",
	AP_DriveStateUnmounting:  "Drive is unmounting",
	AP_DriveStateEmptyClosed: "Drive is empty and closed",
	AP_DriveStateEmptyOpen:   "Drive is empty and open",
	AP_DriveStateInserted:    "Drive has disc inserted",
	AP_DriveStateLoading:     "Drive is loading",
}

func GetDriveStateDescription(id uint) string {
	if desc, ok := driveStateDescriptions[id]; ok {
		return desc
	}
	return "Unknown"
}

const (
	AP_MaxCdromDevices              int = 16
	AP_Progress_MaxValue            int = 65536
	AP_Progress_MaxLayoutItems      int = 10
	AP_UIMSG_BOX_MASK               int = 3854
	AP_UIMSG_BOXOK                  int = 260
	AP_UIMSG_BOXERROR               int = 516
	AP_UIMSG_BOXWARNING             int = 1028
	AP_UIMSG_BOXYESNO               int = 776
	AP_UIMSG_BOXYESNO_ERR           int = 1288
	AP_UIMSG_BOXYESNO_REG           int = 1544
	AP_UIMSG_YES                    int = 0
	AP_UIMSG_NO                     int = 1
	AP_UIMSG_DEBUG                  int = 32
	AP_UIMSG_HIDDEN                 int = 64
	AP_UIMSG_EVENT                  int = 128
	AP_UIMSG_HAVE_URL               int = 131072
	AP_UIMSG_VITEM_BASE             int = 5200
	AP_MMBD_DISC_FLAG_BUSENC        int = 2
	AP_MMBD_MMBD_DISC_FLAG_AACS     int = 4
	AP_MMBD_MMBD_DISC_FLAG_BDPLUS   int = 8
	AP_vastr_Name                   int = 0
	AP_vastr_Version                int = 1
	AP_vastr_Platform               int = 2
	AP_vastr_Build                  int = 3
	AP_vastr_KeyType                int = 4
	AP_vastr_KeyFeatures            int = 5
	AP_vastr_KeyExpiration          int = 6
	AP_vastr_EvalState              int = 7
	AP_vastr_ProgExpiration         int = 8
	AP_vastr_LatestVersion          int = 9
	AP_vastr_RestartRequired        int = 10
	AP_vastr_ExpertMode             int = 11
	AP_vastr_ProfileCount           int = 12
	AP_vastr_ProgExpired            int = 13
	AP_vastr_OutputFolderName       int = 14
	AP_vastr_OutputBaseName         int = 15
	AP_vastr_CurrentProfile         int = 16
	AP_vastr_OpenFileFilter         int = 17
	AP_vastr_WebSiteURL             int = 18
	AP_vastr_OpenDVDFileFilter      int = 19
	AP_vastr_DefaultSelectionString int = 20
	AP_vastr_DefaultOutputFileName  int = 21
	AP_vastr_ExternalAppItem        int = 22
	AP_vastr_InterfaceLanguage      int = 23
	AP_vastr_ProfileString          int = 24
	AP_vastr_KeyString              int = 25
)

type AppItemAttributeId = int

const (
	Unknown                      AppItemAttributeId = 0
	Type                         AppItemAttributeId = 1
	Name                         AppItemAttributeId = 2
	LangCode                     AppItemAttributeId = 3
	LangName                     AppItemAttributeId = 4
	CodecId                      AppItemAttributeId = 5
	CodecShort                   AppItemAttributeId = 6
	CodecLong                    AppItemAttributeId = 7
	ChapterCount                 AppItemAttributeId = 8
	Duration                     AppItemAttributeId = 9
	DiskSize                     AppItemAttributeId = 10
	DiskSizeBytes                AppItemAttributeId = 11
	StreamTypeExtension          AppItemAttributeId = 12
	Bitrate                      AppItemAttributeId = 13
	AudioChannelsCount           AppItemAttributeId = 14
	AngleInfo                    AppItemAttributeId = 15
	SourceFileName               AppItemAttributeId = 16
	AudioSampleRate              AppItemAttributeId = 17
	AudioSampleSize              AppItemAttributeId = 18
	VideoSize                    AppItemAttributeId = 19
	VideoAspectRatio             AppItemAttributeId = 20
	VideoFrameRate               AppItemAttributeId = 21
	StreamFlags                  AppItemAttributeId = 22
	DateTime                     AppItemAttributeId = 23
	OriginalTitleId              AppItemAttributeId = 24
	SegmentsCount                AppItemAttributeId = 25
	SegmentsMap                  AppItemAttributeId = 26
	OutputFileName               AppItemAttributeId = 27
	MetadataLanguageCode         AppItemAttributeId = 28
	MetadataLanguageName         AppItemAttributeId = 29
	TreeInfo                     AppItemAttributeId = 30
	PanelTitle                   AppItemAttributeId = 31
	VolumeName                   AppItemAttributeId = 32
	OrderWeight                  AppItemAttributeId = 33
	OutputFormat                 AppItemAttributeId = 34
	OutputFormatDescription      AppItemAttributeId = 35
	SeamlessInfo                 AppItemAttributeId = 36
	PanelText                    AppItemAttributeId = 37
	MkvFlags                     AppItemAttributeId = 38
	MkvFlagsText                 AppItemAttributeId = 39
	AudioChannelLayoutName       AppItemAttributeId = 40
	OutputCodecShort             AppItemAttributeId = 41
	OutputConversionType         AppItemAttributeId = 42
	OutputAudioSampleRate        AppItemAttributeId = 43
	OutputAudioSampleSize        AppItemAttributeId = 44
	OutputAudioChannelsCount     AppItemAttributeId = 45
	OutputAudioChannelLayoutName AppItemAttributeId = 46
	OutputAudioChannelLayout     AppItemAttributeId = 47
	OutputAudioMixDescription    AppItemAttributeId = 48
	Comment                      AppItemAttributeId = 49
	OffsetSequenceId             AppItemAttributeId = 50
	MaxValue                     AppItemAttributeId = 51
)

const (
	APNotifyUpdateLayoutFlagNoTime   = 1
	APProgressCurrentIndexSourceName = 65280
	APBackupFlagDecryptVideo         = 1
	APOpenFlagManualMode             = 1
	APUpdateDrivesFlagNoScan         = 1
	APUpdateDrivesFlagNoSingleDrive  = 2
)

const (
	APAVStreamFlagDirectorsComments          = 1
	APAVStreamFlagAlternateDirectorsComments = 2
	APAVStreamFlagForVisuallyImpaired        = 4
	APAVStreamFlagCoreAudio                  = 256
	APAVStreamFlagSecondaryAudio             = 512
	APAVStreamFlagHasCoreAudio               = 1024
	APAVStreamFlagDerivedStream              = 2048
	APAVStreamFlagForcedSubtitles            = 4096
	APAVStreamFlagProfileSecondaryStream     = 16384
	APAVStreamFlagOffsetSequenceIdPresent    = 32768
)

const (
	APAPPLOCMAX = 7000
)

const (
	AppDumpDonePartial                        int = 5004
	AppDumpDone                               int = 5005
	AppInitFailed                             int = 5009
	AppAskFolderCreate                        int = 5013
	AppFolderInvalid                          int = 5016
	ProgressAppSaveMkvFreeSpace               int = 5033
	ProtDemoKeyExpired                        int = 5021
	AppKeytypeInvalid                         int = 5095
	AppEvalTimeNever                          int = 5067
	AppBackupFailed                           int = 5069
	AppBackupCompleted                        int = 5070
	AppBackupCompletedHashfail                int = 5079
	ProfileNameDefault                        int = 5086
	VitemName                                 int = 5202
	VitemTimestamp                            int = 5223
	AppIfaceTitle                             int = 6000
	AppCaptionMsg                             int = 6001
	AppAboutboxTitle                          int = 6002
	AppIfaceOpenfileTitle                     int = 6003
	AppSettingdlgTitle                        int = 6135
	AppBackupdlgTitle                         int = 6136
	AppIfaceOpenfileFilterTemplate1           int = 6007
	AppIfaceOpenfileFilterTemplate2           int = 6008
	AppIfaceOpenfolderTitle                   int = 6005
	AppIfaceOpenfolderInfoTitle               int = 6006
	AppIfaceProgressTitle                     int = 6038
	AppIfaceProgressElapsedOnly               int = 6039
	AppIfaceProgressElapsedEta                int = 6040
	AppIfaceActOpenfilesName                  int = 6010
	AppIfaceActOpenfilesSkey                  int = 6011
	AppIfaceActOpenfilesStip                  int = 6012
	AppIfaceActOpenfilesDvdName               int = 6024
	AppIfaceActOpenfilesDvdStip               int = 6026
	AppIfaceActClosediskName                  int = 6013
	AppIfaceActClosediskStip                  int = 6014
	AppIfaceActSetfolderName                  int = 6015
	AppIfaceActSetfolderStip                  int = 6016
	AppIfaceActSaveallmkvName                 int = 6017
	AppIfaceActSaveallmkvStip                 int = 6018
	AppIfaceActCancelName                     int = 6036
	AppIfaceActCancelStip                     int = 6037
	AppIfaceActStreamingName                  int = 6131
	AppIfaceActStreamingStip                  int = 6132
	AppIfaceActBackupName                     int = 6133
	AppIfaceActBackupStip                     int = 6134
	AppIfaceActQuitName                       int = 6019
	AppIfaceActQuitSkey                       int = 6020
	AppIfaceActQuitStip                       int = 6021
	AppIfaceActAboutName                      int = 6022
	AppIfaceActAboutStip                      int = 6023
	AppIfaceActSettingsName                   int = 6042
	AppIfaceActSettingsStip                   int = 6043
	AppIfaceActHelppageName                   int = 6045
	AppIfaceActHelppageStip                   int = 6046
	AppIfaceActRegisterName                   int = 6047
	AppIfaceActRegisterStip                   int = 6048
	AppIfaceActPurchaseName                   int = 6145
	AppIfaceActPurchaseStip                   int = 6146
	AppIfaceActClearlogName                   int = 6110
	AppIfaceActClearlogStip                   int = 6111
	AppIfaceActEjectName                      int = 6052
	AppIfaceActEjectStip                      int = 6053
	AppIfaceActRevertName                     int = 6105
	AppIfaceActRevertStip                     int = 6106
	AppIfaceActNewinstanceName                int = 6107
	AppIfaceActNewinstanceStip                int = 6108
	AppIfaceActOpendiscDvd                    int = 6062
	AppIfaceActOpendiscHddvd                  int = 6063
	AppIfaceActOpendiscBray                   int = 6064
	AppIfaceActOpendiscLoading                int = 6065
	AppIfaceActOpendiscUnknown                int = 6099
	AppIfaceActOpendiscNodisc                 int = 6109
	AppIfaceActTtreeToggle                    int = 6066
	AppIfaceActTtreeSelectAll                 int = 6067
	AppIfaceActTtreeUnselectAll               int = 6068
	AppIfaceMenuFile                          int = 6030
	AppIfaceMenuView                          int = 6031
	AppIfaceMenuHelp                          int = 6032
	AppIfaceMenuToolbar                       int = 6034
	AppIfaceMenuSettings                      int = 6044
	AppIfaceMenuDrives                        int = 6035
	AppIfaceCancelConfirm                     int = 6041
	AppIfaceFatalComm                         int = 6050
	AppIfaceFatalMem                          int = 6051
	AppIfaceGuiVersion                        int = 6054
	AppIfaceLatestVersion                     int = 6158
	AppIfaceLicenseType                       int = 6055
	AppIfaceEvalState                         int = 6056
	AppIfaceEvalExpiration                    int = 6057
	AppIfaceProgExpiration                    int = 6142
	AppIfaceWebsiteUrl                        int = 6159
	AppIfaceVideoFolderNameWin                int = 6058
	AppIfaceVideoFolderNameMac                int = 6059
	AppIfaceVideoFolderNameLinux              int = 6060
	AppIfaceDefaultFolderName                 int = 6061
	AppIfaceMainFrameInfo                     int = 6069
	AppIfaceMainFrameMakeMkv                  int = 6070
	AppIfaceMainFrameProfile                  int = 6180
	AppIfaceMainFrameProperties               int = 6181
	AppIfaceEmptyFrameInfo                    int = 6075
	AppIfaceEmptyFrameSource                  int = 6071
	AppIfaceEmptyFrameType                    int = 6072
	AppIfaceEmptyFrameLabel                   int = 6073
	AppIfaceEmptyFrameProtection              int = 6074
	AppIfaceEmptyFrameDvdManual               int = 6084
	AppIfaceRegisterText                      int = 6076
	AppIfaceRegisterCodeIncorrect             int = 6077
	AppIfaceRegisterCodeNotSaved              int = 6078
	AppIfaceRegisterCodeSaved                 int = 6079
	AppIfaceSettingsIoOptions                 int = 6080
	AppIfaceSettingsIoAuto                    int = 6081
	AppIfaceSettingsIoReadRetry               int = 6082
	AppIfaceSettingsIoReadBuffer              int = 6083
	AppIfaceSettingsIoNoDirectAccess          int = 6150
	AppIfaceSettingsIoDarwinK2Workaround      int = 6151
	AppIfaceSettingsIoSingleDrive             int = 6168
	AppIfaceSettingsDvdAuto                   int = 6085
	AppIfaceSettingsDvdMinLength              int = 6086
	AppIfaceSettingsDvdSpRemove               int = 6087
	AppIfaceSettingsAacsKeyDir                int = 6088
	AppIfaceSettingsBdpMisc                   int = 6129
	AppIfaceSettingsBdpDumpAlways             int = 6130
	AppIfaceSettingsDestTypeNone              int = 6089
	AppIfaceSettingsDestTypeAuto              int = 6090
	AppIfaceSettingsDestTypeSemiauto          int = 6091
	AppIfaceSettingsDestTypeCustom            int = 6092
	AppIfaceSettingsDestdir                   int = 6093
	AppIfaceSettingsGeneralMisc               int = 6094
	AppIfaceSettingsLogDebugMsg               int = 6095
	AppIfaceSettingsDataDir                   int = 6167
	AppIfaceSettingsExpertMode                int = 6169
	AppIfaceSettingsShowAvsync                int = 6170
	AppIfaceSettingsGeneralOnlineUpdates      int = 6188
	AppIfaceSettingsEnableInternetAccess      int = 6187
	AppIfaceSettingsProxyServer               int = 6189
	AppIfaceSettingsTabGeneral                int = 6096
	AppIfaceSettingsMsgFailed                 int = 6097
	AppIfaceSettingsMsgRestart                int = 6098
	AppIfaceSettingsTabLanguage               int = 6152
	AppIfaceSettingsLangInterface             int = 6153
	AppIfaceSettingsLangPreferred             int = 6154
	AppIfaceSettingsLanguageAuto              int = 6156
	AppIfaceSettingsLanguageNone              int = 6157
	AppIfaceSettingsTabIo                     int = 6164
	AppIfaceSettingsTabStreaming              int = 6165
	AppIfaceSettingsTabProt                   int = 6166
	AppIfaceSettingsTabAdvanced               int = 6172
	AppIfaceSettingsAdvDefaultProfile         int = 6173
	AppIfaceSettingsAdvDefaultSelection       int = 6174
	AppIfaceSettingsAdvExternExecPath         int = 6175
	AppIfaceSettingsProtJavaPath              int = 6177
	AppIfaceSettingsAdvOutputFileNameTemplate int = 6178
	AppIfaceSettingsTabIntegration            int = 6190
	AppIfaceSettingsIntText                   int = 6191
	AppIfaceSettingsIntHdrPath                int = 6192
	AppIfaceKeyText                           int = 6179
	AppIfaceKeyName                           int = 6182
	AppIfaceKeyType                           int = 6183
	AppIfaceKeyDate                           int = 6184
	AppIfaceBackupdlgTextCaption              int = 6137
	AppIfaceBackupdlgText                     int = 6138
	AppIfaceBackupdlgFolder                   int = 6139
	AppIfaceBackupdlgOptions                  int = 6147
	AppIfaceBackupdlgDecrypt                  int = 6148
	AppIfaceDriveinfoLoading                  int = 6100
	AppIfaceDriveinfoUnmounting               int = 6112
	AppIfaceDriveinfoWait                     int = 6101
	AppIfaceDriveinfoNodisc                   int = 6102
	AppIfaceDriveinfoDatadisc                 int = 6103
	AppIfaceDriveinfoNone                     int = 6104
	AppIfaceFlagsDirectorsComments            int = 6125
	AppIfaceFlagsAltDirectorsComments         int = 6126
	AppIfaceFlagsSecondaryAudio               int = 6127
	AppIfaceFlagsForVisuallyImpaired          int = 6128
	AppIfaceFlagsCoreAudio                    int = 6143
	AppIfaceFlagsForcedSubtitles              int = 6144
	AppIfaceFlagsProfileSecondaryStream       int = 6171
	AppIfaceIteminfoSource                    int = 6119
	AppIfaceIteminfoTitle                     int = 6120
	AppIfaceIteminfoTrack                     int = 6121
	AppIfaceIteminfoAttachment                int = 6122
	AppIfaceIteminfoChapter                   int = 6123
	AppIfaceIteminfoChapters                  int = 6124
	AppTtreeTitle                             int = 6200
	AppTtreeVideo                             int = 6201
	AppTtreeAudio                             int = 6202
	AppTtreeSubpicture                        int = 6203
	AppTtreeAttachment                        int = 6214
	AppTtreeChapters                          int = 6215
	AppTtreeChapter                           int = 6216
	AppTtreeForcedSubtitles                   int = 6211
	AppTtreeHdrType                           int = 6204
	AppTtreeHdrDesc                           int = 6205
	DvdTypeDisk                               int = 6206
	BrayTypeDisk                              int = 6209
	HddvdTypeDisk                             int = 6212
	MkvTypeFile                               int = 6213
	AppTtreeChapDesc                          int = 6207
	AppTtreeAngleDesc                         int = 6210
	AppDvdManualTitle                         int = 6220
	AppDvdManualText                          int = 6225
	AppDvdTitlesCount                         int = 6221
	AppDvdCountCells                          int = 6222
	AppDvdCountPgc                            int = 6223
	AppDvdBrokenTitleEntry                    int = 6224
	AppSingleDriveTitle                       int = 6226
	AppSingleDriveText                        int = 6227
	AppSingleDriveAll                         int = 6228
	AppSingleDriveCaption                     int = 6229
	AppSiDriveinfo                            int = 6300
	AppSiProfile                              int = 6301
	AppSiManufacturer                         int = 6302
	AppSiProduct                              int = 6303
	AppSiRevision                             int = 6304
	AppSiSerial                               int = 6305
	AppSiFirmware                             int = 6306
	AppSiFirdate                              int = 6307
	AppSiBecflags                             int = 6308
	AppSiHighestAacs                          int = 6309
	AppSiDiscinfo                             int = 6320
	AppSiNodisc                               int = 6321
	AppSiDiscload                             int = 6322
	AppSiCapacity                             int = 6323
	AppSiDisctype                             int = 6324
	AppSiDiscsize                             int = 6325
	AppSiDiscrate                             int = 6326
	AppSiDisclayers                           int = 6327
	AppSiDisccbl                              int = 6329
	AppSiDisccbl25                            int = 6330
	AppSiDisccbl27                            int = 6331
	AppSiDevice                               int = 6332
)

var appConstantsDescriptions = map[int]string{
	AppDumpDonePartial:                        "Partial dump done",
	AppDumpDone:                               "Dump done",
	AppInitFailed:                             "Initialization failed",
	AppAskFolderCreate:                        "Ask to create folder",
	AppFolderInvalid:                          "Invalid folder",
	ProgressAppSaveMkvFreeSpace:               "Save MKV free space progress",
	ProtDemoKeyExpired:                        "Demo key expired",
	AppKeytypeInvalid:                         "Invalid key type",
	AppEvalTimeNever:                          "Evaluation time never",
	AppBackupFailed:                           "Backup failed",
	AppBackupCompleted:                        "Backup completed",
	AppBackupCompletedHashfail:                "Backup completed with hash failure",
	ProfileNameDefault:                        "Default profile name",
	VitemName:                                 "Virtual item name",
	VitemTimestamp:                            "Virtual item timestamp",
	AppIfaceTitle:                             "Interface title",
	AppCaptionMsg:                             "Caption message",
	AppAboutboxTitle:                          "About box title",
	AppIfaceOpenfileTitle:                     "Open file title",
	AppSettingdlgTitle:                        "Settings dialog title",
	AppBackupdlgTitle:                         "Backup dialog title",
	AppIfaceOpenfileFilterTemplate1:           "Open file filter template 1",
	AppIfaceOpenfileFilterTemplate2:           "Open file filter template 2",
	AppIfaceOpenfolderTitle:                   "Open folder title",
	AppIfaceOpenfolderInfoTitle:               "Open folder info title",
	AppIfaceProgressTitle:                     "Progress title",
	AppIfaceProgressElapsedOnly:               "Progress elapsed only",
	AppIfaceProgressElapsedEta:                "Progress elapsed and ETA",
	AppIfaceActOpenfilesName:                  "Open files action name",
	AppIfaceActOpenfilesSkey:                  "Open files shortcut key",
	AppIfaceActOpenfilesStip:                  "Open files tooltip",
	AppIfaceActOpenfilesDvdName:               "Open DVD files action name",
	AppIfaceActOpenfilesDvdStip:               "Open DVD files tooltip",
	AppIfaceActClosediskName:                  "Close disk action name",
	AppIfaceActClosediskStip:                  "Close disk tooltip",
	AppIfaceActSetfolderName:                  "Set folder action name",
	AppIfaceActSetfolderStip:                  "Set folder tooltip",
	AppIfaceActSaveallmkvName:                 "Save all MKV action name",
	AppIfaceActSaveallmkvStip:                 "Save all MKV tooltip",
	AppIfaceActCancelName:                     "Cancel action name",
	AppIfaceActCancelStip:                     "Cancel tooltip",
	AppIfaceActStreamingName:                  "Streaming action name",
	AppIfaceActStreamingStip:                  "Streaming tooltip",
	AppIfaceActBackupName:                     "Backup action name",
	AppIfaceActBackupStip:                     "Backup tooltip",
	AppIfaceActQuitName:                       "Quit action name",
	AppIfaceActQuitSkey:                       "Quit shortcut key",
	AppIfaceActQuitStip:                       "Quit tooltip",
	AppIfaceActAboutName:                      "About action name",
	AppIfaceActAboutStip:                      "About tooltip",
	AppIfaceActSettingsName:                   "Settings action name",
	AppIfaceActSettingsStip:                   "Settings tooltip",
	AppIfaceActHelppageName:                   "Help page action name",
	AppIfaceActHelppageStip:                   "Help page tooltip",
	AppIfaceActRegisterName:                   "Register action name",
	AppIfaceActRegisterStip:                   "Register tooltip",
	AppIfaceActPurchaseName:                   "Purchase action name",
	AppIfaceActPurchaseStip:                   "Purchase tooltip",
	AppIfaceActClearlogName:                   "Clear log action name",
	AppIfaceActClearlogStip:                   "Clear log tooltip",
	AppIfaceActEjectName:                      "Eject action name",
	AppIfaceActEjectStip:                      "Eject tooltip",
	AppIfaceActRevertName:                     "Revert action name",
	AppIfaceActRevertStip:                     "Revert tooltip",
	AppIfaceActNewinstanceName:                "New instance action name",
	AppIfaceActNewinstanceStip:                "New instance tooltip",
	AppIfaceActOpendiscDvd:                    "Open DVD disc action",
	AppIfaceActOpendiscHddvd:                  "Open HD DVD disc action",
	AppIfaceActOpendiscBray:                   "Open Blu-ray disc action",
	AppIfaceActOpendiscLoading:                "Open disc loading action",
	AppIfaceActOpendiscUnknown:                "Open unknown disc action",
	AppIfaceActOpendiscNodisc:                 "Open no disc action",
	AppIfaceActTtreeToggle:                    "Toggle tree action",
	AppIfaceActTtreeSelectAll:                 "Select all tree action",
	AppIfaceActTtreeUnselectAll:               "Unselect all tree action",
	AppIfaceMenuFile:                          "File menu",
	AppIfaceMenuView:                          "View menu",
	AppIfaceMenuHelp:                          "Help menu",
	AppIfaceMenuToolbar:                       "Toolbar menu",
	AppIfaceMenuSettings:                      "Settings menu",
	AppIfaceMenuDrives:                        "Drives menu",
	AppIfaceCancelConfirm:                     "Cancel confirmation",
	AppIfaceFatalComm:                         "Fatal communication error",
	AppIfaceFatalMem:                          "Fatal memory error",
	AppIfaceGuiVersion:                        "GUI version",
	AppIfaceLatestVersion:                     "Latest version",
	AppIfaceLicenseType:                       "License type",
	AppIfaceEvalState:                         "Evaluation state",
	AppIfaceEvalExpiration:                    "Evaluation expiration",
	AppIfaceProgExpiration:                    "Program expiration",
	AppIfaceWebsiteUrl:                        "Website URL",
	AppIfaceVideoFolderNameWin:                "Video folder name (Windows)",
	AppIfaceVideoFolderNameMac:                "Video folder name (Mac)",
	AppIfaceVideoFolderNameLinux:              "Video folder name (Linux)",
	AppIfaceDefaultFolderName:                 "Default folder name",
	AppIfaceMainFrameInfo:                     "Main frame info",
	AppIfaceMainFrameMakeMkv:                  "Main frame MakeMKV",
	AppIfaceMainFrameProfile:                  "Main frame profile",
	AppIfaceMainFrameProperties:               "Main frame properties",
	AppIfaceEmptyFrameInfo:                    "Empty frame info",
	AppIfaceEmptyFrameSource:                  "Empty frame source",
	AppIfaceEmptyFrameType:                    "Empty frame type",
	AppIfaceEmptyFrameLabel:                   "Empty frame label",
	AppIfaceEmptyFrameProtection:              "Empty frame protection",
	AppIfaceEmptyFrameDvdManual:               "Empty frame DVD manual",
	AppIfaceRegisterText:                      "Register text",
	AppIfaceRegisterCodeIncorrect:             "Register code incorrect",
	AppIfaceRegisterCodeNotSaved:              "Register code not saved",
	AppIfaceRegisterCodeSaved:                 "Register code saved",
	AppIfaceSettingsIoOptions:                 "Settings IO options",
	AppIfaceSettingsIoAuto:                    "Settings IO auto",
	AppIfaceSettingsIoReadRetry:               "Settings IO read retry",
	AppIfaceSettingsIoReadBuffer:              "Settings IO read buffer",
	AppIfaceSettingsIoNoDirectAccess:          "Settings IO no direct access",
	AppIfaceSettingsIoDarwinK2Workaround:      "Settings IO Darwin K2 workaround",
	AppIfaceSettingsIoSingleDrive:             "Settings IO single drive",
	AppIfaceSettingsDvdAuto:                   "Settings DVD auto",
	AppIfaceSettingsDvdMinLength:              "Settings DVD minimum length",
	AppIfaceSettingsDvdSpRemove:               "Settings DVD SP remove",
	AppIfaceSettingsAacsKeyDir:                "Settings AACS key directory",
	AppIfaceSettingsBdpMisc:                   "Settings BDP miscellaneous",
	AppIfaceSettingsBdpDumpAlways:             "Settings BDP dump always",
	AppIfaceSettingsDestTypeNone:              "Settings destination type none",
	AppIfaceSettingsDestTypeAuto:              "Settings destination type auto",
	AppIfaceSettingsDestTypeSemiauto:          "Settings destination type semi-auto",
	AppIfaceSettingsDestTypeCustom:            "Settings destination type custom",
	AppIfaceSettingsDestdir:                   "Settings destination directory",
	AppIfaceSettingsGeneralMisc:               "Settings general miscellaneous",
	AppIfaceSettingsLogDebugMsg:               "Settings log debug messages",
	AppIfaceSettingsDataDir:                   "Settings data directory",
	AppIfaceSettingsExpertMode:                "Settings expert mode",
	AppIfaceSettingsShowAvsync:                "Settings show AV sync",
	AppIfaceSettingsGeneralOnlineUpdates:      "Settings general online updates",
	AppIfaceSettingsEnableInternetAccess:      "Settings enable internet access",
	AppIfaceSettingsProxyServer:               "Settings proxy server",
	AppIfaceSettingsTabGeneral:                "Settings tab general",
	AppIfaceSettingsMsgFailed:                 "Settings message failed",
	AppIfaceSettingsMsgRestart:                "Settings message restart",
	AppIfaceSettingsTabLanguage:               "Settings tab language",
	AppIfaceSettingsLangInterface:             "Settings language interface",
	AppIfaceSettingsLangPreferred:             "Settings language preferred",
	AppIfaceSettingsLanguageAuto:              "Settings language auto",
	AppIfaceSettingsLanguageNone:              "Settings language none",
	AppIfaceSettingsTabIo:                     "Settings tab IO",
	AppIfaceSettingsTabStreaming:              "Settings tab streaming",
	AppIfaceSettingsTabProt:                   "Settings tab protection",
	AppIfaceSettingsTabAdvanced:               "Settings tab advanced",
	AppIfaceSettingsAdvDefaultProfile:         "Settings advanced default profile",
	AppIfaceSettingsAdvDefaultSelection:       "Settings advanced default selection",
	AppIfaceSettingsAdvExternExecPath:         "Settings advanced external execution path",
	AppIfaceSettingsProtJavaPath:              "Settings protection Java path",
	AppIfaceSettingsAdvOutputFileNameTemplate: "Settings advanced output file name template",
	AppIfaceSettingsTabIntegration:            "Settings tab integration",
	AppIfaceSettingsIntText:                   "Settings integration text",
	AppIfaceSettingsIntHdrPath:                "Settings integration HDR path",
	AppIfaceKeyText:                           "Key text",
	AppIfaceKeyName:                           "Key name",
	AppIfaceKeyType:                           "Key type",
	AppIfaceKeyDate:                           "Key date",
	AppIfaceBackupdlgTextCaption:              "Backup dialog text caption",
	AppIfaceBackupdlgText:                     "Backup dialog text",
	AppIfaceBackupdlgFolder:                   "Backup dialog folder",
	AppIfaceBackupdlgOptions:                  "Backup dialog options",
	AppIfaceBackupdlgDecrypt:                  "Backup dialog decrypt",
	AppIfaceDriveinfoLoading:                  "Drive info loading",
	AppIfaceDriveinfoUnmounting:               "Drive info unmounting",
	AppIfaceDriveinfoWait:                     "Drive info wait",
	AppIfaceDriveinfoNodisc:                   "Drive info no disc",
	AppIfaceDriveinfoDatadisc:                 "Drive info data disc",
	AppIfaceDriveinfoNone:                     "Drive info none",
	AppIfaceFlagsDirectorsComments:            "Flags director's comments",
	AppIfaceFlagsAltDirectorsComments:         "Flags alternate director's comments",
	AppIfaceFlagsSecondaryAudio:               "Flags secondary audio",
	AppIfaceFlagsForVisuallyImpaired:          "Flags for visually impaired",
	AppIfaceFlagsCoreAudio:                    "Flags core audio",
	AppIfaceFlagsForcedSubtitles:              "Flags forced subtitles",
	AppIfaceFlagsProfileSecondaryStream:       "Flags profile secondary stream",
	AppIfaceIteminfoSource:                    "Item info source",
	AppIfaceIteminfoTitle:                     "Item info title",
	AppIfaceIteminfoTrack:                     "Item info track",
	AppIfaceIteminfoAttachment:                "Item info attachment",
	AppIfaceIteminfoChapter:                   "Item info chapter",
	AppIfaceIteminfoChapters:                  "Item info chapters",
	AppTtreeTitle:                             "Tree title",
	AppTtreeVideo:                             "Tree video",
	AppTtreeAudio:                             "Tree audio",
	AppTtreeSubpicture:                        "Tree subpicture",
	AppTtreeAttachment:                        "Tree attachment",
	AppTtreeChapters:                          "Tree chapters",
	AppTtreeChapter:                           "Tree chapter",
	AppTtreeForcedSubtitles:                   "Tree forced subtitles",
	AppTtreeHdrType:                           "Tree HDR type",
	AppTtreeHdrDesc:                           "Tree HDR description",
	DvdTypeDisk:                               "DVD type disk",
	BrayTypeDisk:                              "Blu-ray type disk",
	HddvdTypeDisk:                             "HD DVD type disk",
	MkvTypeFile:                               "MKV type file",
	AppTtreeChapDesc:                          "Tree chapter description",
	AppTtreeAngleDesc:                         "Tree angle description",
	AppDvdManualTitle:                         "DVD manual title",
	AppDvdManualText:                          "DVD manual text",
	AppDvdTitlesCount:                         "DVD titles count",
	AppDvdCountCells:                          "DVD count cells",
	AppDvdCountPgc:                            "DVD count PGC",
	AppDvdBrokenTitleEntry:                    "DVD broken title entry",
	AppSingleDriveTitle:                       "Single drive title",
	AppSingleDriveText:                        "Single drive text",
	AppSingleDriveAll:                         "Single drive all",
	AppSingleDriveCaption:                     "Single drive caption",
	AppSiDriveinfo:                            "SI drive info",
	AppSiProfile:                              "SI profile",
	AppSiManufacturer:                         "SI manufacturer",
	AppSiProduct:                              "SI product",
	AppSiRevision:                             "SI revision",
	AppSiSerial:                               "SI serial",
	AppSiFirmware:                             "SI firmware",
	AppSiFirdate:                              "SI firmware date",
	AppSiBecflags:                             "SI BEC flags",
	AppSiHighestAacs:                          "SI highest AACS",
	AppSiDiscinfo:                             "SI disc info",
	AppSiNodisc:                               "SI no disc",
	AppSiDiscload:                             "SI disc load",
	AppSiCapacity:                             "SI capacity",
	AppSiDisctype:                             "SI disc type",
	AppSiDiscsize:                             "SI disc size",
	AppSiDiscrate:                             "SI disc rate",
	AppSiDisclayers:                           "SI disc layers",
	AppSiDisccbl:                              "SI disc CBL",
	AppSiDisccbl25:                            "SI disc CBL 25",
	AppSiDisccbl27:                            "SI disc CBL 27",
	AppSiDevice:                               "SI device",
}

func GetAppConstantDescription(id int) (string, error) {
	if desc, ok := appConstantsDescriptions[id]; ok {
		return desc, nil
	}
	return "", errors.New("Unknown")
}
