package ids

type ApItemAttributeId int

const (
	AP_MaxCdromDevices              uint = 16
	AP_Progress_MaxValue            uint = 65536
	AP_Progress_MaxLayoutItems      uint = 10
	AP_UIMSG_BOX_MASK               uint = 3854
	AP_UIMSG_BOXOK                  uint = 260
	AP_UIMSG_BOXERROR               uint = 516
	AP_UIMSG_BOXWARNING             uint = 1028
	AP_UIMSG_BOXYESNO               uint = 776
	AP_UIMSG_BOXYESNO_ERR           uint = 1288
	AP_UIMSG_BOXYESNO_REG           uint = 1544
	AP_UIMSG_YES                    uint = 0
	AP_UIMSG_NO                     uint = 1
	AP_UIMSG_DEBUG                  uint = 32
	AP_UIMSG_HIDDEN                 uint = 64
	AP_UIMSG_EVENT                  uint = 128
	AP_UIMSG_HAVE_URL               uint = 131072
	AP_UIMSG_VITEM_BASE             uint = 5200
	AP_MMBD_DISC_FLAG_BUSENC        uint = 2
	AP_MMBD_MMBD_DISC_FLAG_AACS     uint = 4
	AP_MMBD_MMBD_DISC_FLAG_BDPLUS   uint = 8
	AP_vastr_Name                   uint = 0
	AP_vastr_Version                uint = 1
	AP_vastr_Platform               uint = 2
	AP_vastr_Build                  uint = 3
	AP_vastr_KeyType                uint = 4
	AP_vastr_KeyFeatures            uint = 5
	AP_vastr_KeyExpiration          uint = 6
	AP_vastr_EvalState              uint = 7
	AP_vastr_ProgExpiration         uint = 8
	AP_vastr_LatestVersion          uint = 9
	AP_vastr_RestartRequired        uint = 10
	AP_vastr_ExpertMode             uint = 11
	AP_vastr_ProfileCount           uint = 12
	AP_vastr_ProgExpired            uint = 13
	AP_vastr_OutputFolderName       uint = 14
	AP_vastr_OutputBaseName         uint = 15
	AP_vastr_CurrentProfile         uint = 16
	AP_vastr_OpenFileFilter         uint = 17
	AP_vastr_WebSiteURL             uint = 18
	AP_vastr_OpenDVDFileFilter      uint = 19
	AP_vastr_DefaultSelectionString uint = 20
	AP_vastr_DefaultOutputFileName  uint = 21
	AP_vastr_ExternalAppItem        uint = 22
	AP_vastr_InterfaceLanguage      uint = 23
	AP_vastr_ProfileString          uint = 24
	AP_vastr_KeyString              uint = 25
)

type AppItemAttributeId = uint

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

var attributeDetailedDescription = map[AppItemAttributeId]string{
	Unknown:                      "Unknown",
	Type:                         "Type",
	Name:                         "Name",
	LangCode:                     "Language Code",
	LangName:                     "Language Name",
	CodecId:                      "Codec ID",
	CodecShort:                   "Short Codec Name",
	CodecLong:                    "Long Codec Name",
	ChapterCount:                 "Number of Chapters",
	Duration:                     "Duration",
	DiskSize:                     "Disk Size",
	DiskSizeBytes:                "Disk Size in Bytes",
	StreamTypeExtension:          "Stream Type Extension",
	Bitrate:                      "Bitrate",
	AudioChannelsCount:           "Number of Audio Channels",
	AngleInfo:                    "Angle Information",
	SourceFileName:               "Source File Name",
	AudioSampleRate:              "Audio Sample Rate",
	AudioSampleSize:              "Audio Sample Size",
	VideoSize:                    "Video Size",
	VideoAspectRatio:             "Video Aspect Ratio",
	VideoFrameRate:               "Video Frame Rate",
	StreamFlags:                  "Stream Flags",
	DateTime:                     "Date and Time",
	OriginalTitleId:              "Original Title ID",
	SegmentsCount:                "Number of Segments",
	SegmentsMap:                  "Segments Map",
	OutputFileName:               "Output File Name",
	MetadataLanguageCode:         "Metadata Language Code",
	MetadataLanguageName:         "Metadata Language Name",
	TreeInfo:                     "Tree Information",
	PanelTitle:                   "Panel Title",
	VolumeName:                   "Volume Name",
	OrderWeight:                  "Order Weight",
	OutputFormat:                 "Output Format",
	OutputFormatDescription:      "Output Format Description",
	SeamlessInfo:                 "Seamless Information",
	PanelText:                    "Panel Text",
	MkvFlags:                     "MKV Flags",
	MkvFlagsText:                 "MKV Flags Text",
	AudioChannelLayoutName:       "Audio Channel Layout Name",
	OutputCodecShort:             "Output Short Codec Name",
	OutputConversionType:         "Output Conversion Type",
	OutputAudioSampleRate:        "Output Audio Sample Rate",
	OutputAudioSampleSize:        "Output Audio Sample Size",
	OutputAudioChannelsCount:     "Output Number of Audio Channels",
	OutputAudioChannelLayoutName: "Output Audio Channel Layout Name",
	OutputAudioChannelLayout:     "Output Audio Channel Layout",
	OutputAudioMixDescription:    "Output Audio Mix Description",
	Comment:                      "Comment",
	OffsetSequenceId:             "Offset Sequence ID",
	MaxValue:                     "Maximum Value",
}

func GetItemAttributeDescription(id uint) string {
	if desc, ok := attributeDetailedDescription[id]; ok {
		return desc
	}
	return "Unknown Application Item Attribute"
}

const (
	AP_DskFsFlagDvdFilesPresent    uint = 1
	AP_DskFsFlagHdvdFilesPresent   uint = 2
	AP_DskFsFlagBlurayFilesPresent uint = 4
	AP_DskFsFlagAacsFilesPresent   uint = 8
	AP_DskFsFlagBdsvmFilesPresent  uint = 16
)

func GetDiskFileFlagDescription(id uint) string {
	switch id {
	case AP_DskFsFlagDvdFilesPresent:
		return "DVD files present on disk"
	case AP_DskFsFlagHdvdFilesPresent:
		return "HD DVD files present on disk"
	case AP_DskFsFlagBlurayFilesPresent:
		return "Blu-ray files present on disk"
	case AP_DskFsFlagAacsFilesPresent:
		return "Aacs files present on disk"
	case AP_DskFsFlagBdsvmFilesPresent:
		return "Blu-ray disc movie folder files present on disk"
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

func GetDriveStateDescription(id uint) string {
	switch id {
	case AP_DriveStateNoDrive:
		return "No Drive detected"
	case AP_DriveStateUnmounting:
		return "Drive is unmounting"
	case AP_DriveStateEmptyClosed:
		return "Drive is empty and closed"
	case AP_DriveStateEmptyOpen:
		return "Drive is empty and open"
	case AP_DriveStateInserted:
		return "Drive has disc inserted"
	case AP_DriveStateLoading:
		return "Drive is loading"
	}
	return "Unknown"
}

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
	AppDumpDonePartial                        uint = 5004
	AppDumpDone                               uint = 5005
	AppInitFailed                             uint = 5009
	AppAskFolderCreate                        uint = 5013
	AppFolderInvalid                          uint = 5016
	ProgressAppSaveMkvFreeSpace               uint = 5033
	ProtDemoKeyExpired                        uint = 5021
	AppKeytypeInvalid                         uint = 5095
	AppEvalTimeNever                          uint = 5067
	AppBackupFailed                           uint = 5069
	AppBackupCompleted                        uint = 5070
	AppBackupCompletedHashfail                uint = 5079
	ProfileNameDefault                        uint = 5086
	VitemName                                 uint = 5202
	VitemTimestamp                            uint = 5223
	AppIfaceTitle                             uint = 6000
	AppCaptionMsg                             uint = 6001
	AppAboutboxTitle                          uint = 6002
	AppIfaceOpenfileTitle                     uint = 6003
	AppSettingdlgTitle                        uint = 6135
	AppBackupdlgTitle                         uint = 6136
	AppIfaceOpenfileFilterTemplate1           uint = 6007
	AppIfaceOpenfileFilterTemplate2           uint = 6008
	AppIfaceOpenfolderTitle                   uint = 6005
	AppIfaceOpenfolderInfoTitle               uint = 6006
	AppIfaceProgressTitle                     uint = 6038
	AppIfaceProgressElapsedOnly               uint = 6039
	AppIfaceProgressElapsedEta                uint = 6040
	AppIfaceActOpenfilesName                  uint = 6010
	AppIfaceActOpenfilesSkey                  uint = 6011
	AppIfaceActOpenfilesStip                  uint = 6012
	AppIfaceActOpenfilesDvdName               uint = 6024
	AppIfaceActOpenfilesDvdStip               uint = 6026
	AppIfaceActClosediskName                  uint = 6013
	AppIfaceActClosediskStip                  uint = 6014
	AppIfaceActSetfolderName                  uint = 6015
	AppIfaceActSetfolderStip                  uint = 6016
	AppIfaceActSaveallmkvName                 uint = 6017
	AppIfaceActSaveallmkvStip                 uint = 6018
	AppIfaceActCancelName                     uint = 6036
	AppIfaceActCancelStip                     uint = 6037
	AppIfaceActStreamingName                  uint = 6131
	AppIfaceActStreamingStip                  uint = 6132
	AppIfaceActBackupName                     uint = 6133
	AppIfaceActBackupStip                     uint = 6134
	AppIfaceActQuitName                       uint = 6019
	AppIfaceActQuitSkey                       uint = 6020
	AppIfaceActQuitStip                       uint = 6021
	AppIfaceActAboutName                      uint = 6022
	AppIfaceActAboutStip                      uint = 6023
	AppIfaceActSettingsName                   uint = 6042
	AppIfaceActSettingsStip                   uint = 6043
	AppIfaceActHelppageName                   uint = 6045
	AppIfaceActHelppageStip                   uint = 6046
	AppIfaceActRegisterName                   uint = 6047
	AppIfaceActRegisterStip                   uint = 6048
	AppIfaceActPurchaseName                   uint = 6145
	AppIfaceActPurchaseStip                   uint = 6146
	AppIfaceActClearlogName                   uint = 6110
	AppIfaceActClearlogStip                   uint = 6111
	AppIfaceActEjectName                      uint = 6052
	AppIfaceActEjectStip                      uint = 6053
	AppIfaceActRevertName                     uint = 6105
	AppIfaceActRevertStip                     uint = 6106
	AppIfaceActNewinstanceName                uint = 6107
	AppIfaceActNewinstanceStip                uint = 6108
	AppIfaceActOpendiscDvd                    uint = 6062
	AppIfaceActOpendiscHddvd                  uint = 6063
	AppIfaceActOpendiscBray                   uint = 6064
	AppIfaceActOpendiscLoading                uint = 6065
	AppIfaceActOpendiscUnknown                uint = 6099
	AppIfaceActOpendiscNodisc                 uint = 6109
	AppIfaceActTtreeToggle                    uint = 6066
	AppIfaceActTtreeSelectAll                 uint = 6067
	AppIfaceActTtreeUnselectAll               uint = 6068
	AppIfaceMenuFile                          uint = 6030
	AppIfaceMenuView                          uint = 6031
	AppIfaceMenuHelp                          uint = 6032
	AppIfaceMenuToolbar                       uint = 6034
	AppIfaceMenuSettings                      uint = 6044
	AppIfaceMenuDrives                        uint = 6035
	AppIfaceCancelConfirm                     uint = 6041
	AppIfaceFatalComm                         uint = 6050
	AppIfaceFatalMem                          uint = 6051
	AppIfaceGuiVersion                        uint = 6054
	AppIfaceLatestVersion                     uint = 6158
	AppIfaceLicenseType                       uint = 6055
	AppIfaceEvalState                         uint = 6056
	AppIfaceEvalExpiration                    uint = 6057
	AppIfaceProgExpiration                    uint = 6142
	AppIfaceWebsiteUrl                        uint = 6159
	AppIfaceVideoFolderNameWin                uint = 6058
	AppIfaceVideoFolderNameMac                uint = 6059
	AppIfaceVideoFolderNameLinux              uint = 6060
	AppIfaceDefaultFolderName                 uint = 6061
	AppIfaceMainFrameInfo                     uint = 6069
	AppIfaceMainFrameMakeMkv                  uint = 6070
	AppIfaceMainFrameProfile                  uint = 6180
	AppIfaceMainFrameProperties               uint = 6181
	AppIfaceEmptyFrameInfo                    uint = 6075
	AppIfaceEmptyFrameSource                  uint = 6071
	AppIfaceEmptyFrameType                    uint = 6072
	AppIfaceEmptyFrameLabel                   uint = 6073
	AppIfaceEmptyFrameProtection              uint = 6074
	AppIfaceEmptyFrameDvdManual               uint = 6084
	AppIfaceRegisterText                      uint = 6076
	AppIfaceRegisterCodeIncorrect             uint = 6077
	AppIfaceRegisterCodeNotSaved              uint = 6078
	AppIfaceRegisterCodeSaved                 uint = 6079
	AppIfaceSettingsIoOptions                 uint = 6080
	AppIfaceSettingsIoAuto                    uint = 6081
	AppIfaceSettingsIoReadRetry               uint = 6082
	AppIfaceSettingsIoReadBuffer              uint = 6083
	AppIfaceSettingsIoNoDirectAccess          uint = 6150
	AppIfaceSettingsIoDarwinK2Workaround      uint = 6151
	AppIfaceSettingsIoSingleDrive             uint = 6168
	AppIfaceSettingsDvdAuto                   uint = 6085
	AppIfaceSettingsDvdMinLength              uint = 6086
	AppIfaceSettingsDvdSpRemove               uint = 6087
	AppIfaceSettingsAacsKeyDir                uint = 6088
	AppIfaceSettingsBdpMisc                   uint = 6129
	AppIfaceSettingsBdpDumpAlways             uint = 6130
	AppIfaceSettingsDestTypeNone              uint = 6089
	AppIfaceSettingsDestTypeAuto              uint = 6090
	AppIfaceSettingsDestTypeSemiauto          uint = 6091
	AppIfaceSettingsDestTypeCustom            uint = 6092
	AppIfaceSettingsDestdir                   uint = 6093
	AppIfaceSettingsGeneralMisc               uint = 6094
	AppIfaceSettingsLogDebugMsg               uint = 6095
	AppIfaceSettingsDataDir                   uint = 6167
	AppIfaceSettingsExpertMode                uint = 6169
	AppIfaceSettingsShowAvsync                uint = 6170
	AppIfaceSettingsGeneralOnlineUpdates      uint = 6188
	AppIfaceSettingsEnableInternetAccess      uint = 6187
	AppIfaceSettingsProxyServer               uint = 6189
	AppIfaceSettingsTabGeneral                uint = 6096
	AppIfaceSettingsMsgFailed                 uint = 6097
	AppIfaceSettingsMsgRestart                uint = 6098
	AppIfaceSettingsTabLanguage               uint = 6152
	AppIfaceSettingsLangInterface             uint = 6153
	AppIfaceSettingsLangPreferred             uint = 6154
	AppIfaceSettingsLanguageAuto              uint = 6156
	AppIfaceSettingsLanguageNone              uint = 6157
	AppIfaceSettingsTabIo                     uint = 6164
	AppIfaceSettingsTabStreaming              uint = 6165
	AppIfaceSettingsTabProt                   uint = 6166
	AppIfaceSettingsTabAdvanced               uint = 6172
	AppIfaceSettingsAdvDefaultProfile         uint = 6173
	AppIfaceSettingsAdvDefaultSelection       uint = 6174
	AppIfaceSettingsAdvExternExecPath         uint = 6175
	AppIfaceSettingsProtJavaPath              uint = 6177
	AppIfaceSettingsAdvOutputFileNameTemplate uint = 6178
	AppIfaceSettingsTabIntegration            uint = 6190
	AppIfaceSettingsIntText                   uint = 6191
	AppIfaceSettingsIntHdrPath                uint = 6192
	AppIfaceKeyText                           uint = 6179
	AppIfaceKeyName                           uint = 6182
	AppIfaceKeyType                           uint = 6183
	AppIfaceKeyDate                           uint = 6184
	AppIfaceBackupdlgTextCaption              uint = 6137
	AppIfaceBackupdlgText                     uint = 6138
	AppIfaceBackupdlgFolder                   uint = 6139
	AppIfaceBackupdlgOptions                  uint = 6147
	AppIfaceBackupdlgDecrypt                  uint = 6148
	AppIfaceDriveinfoLoading                  uint = 6100
	AppIfaceDriveinfoUnmounting               uint = 6112
	AppIfaceDriveinfoWait                     uint = 6101
	AppIfaceDriveinfoNodisc                   uint = 6102
	AppIfaceDriveinfoDatadisc                 uint = 6103
	AppIfaceDriveinfoNone                     uint = 6104
	AppIfaceFlagsDirectorsComments            uint = 6125
	AppIfaceFlagsAltDirectorsComments         uint = 6126
	AppIfaceFlagsSecondaryAudio               uint = 6127
	AppIfaceFlagsForVisuallyImpaired          uint = 6128
	AppIfaceFlagsCoreAudio                    uint = 6143
	AppIfaceFlagsForcedSubtitles              uint = 6144
	AppIfaceFlagsProfileSecondaryStream       uint = 6171
	AppIfaceIteminfoSource                    uint = 6119
	AppIfaceIteminfoTitle                     uint = 6120
	AppIfaceIteminfoTrack                     uint = 6121
	AppIfaceIteminfoAttachment                uint = 6122
	AppIfaceIteminfoChapter                   uint = 6123
	AppIfaceIteminfoChapters                  uint = 6124
	AppTtreeTitle                             uint = 6200
	AppTtreeVideo                             uint = 6201
	AppTtreeAudio                             uint = 6202
	AppTtreeSubpicture                        uint = 6203
	AppTtreeAttachment                        uint = 6214
	AppTtreeChapters                          uint = 6215
	AppTtreeChapter                           uint = 6216
	AppTtreeForcedSubtitles                   uint = 6211
	AppTtreeHdrType                           uint = 6204
	AppTtreeHdrDesc                           uint = 6205
	DvdTypeDisk                               uint = 6206
	BrayTypeDisk                              uint = 6209
	HddvdTypeDisk                             uint = 6212
	MkvTypeFile                               uint = 6213
	AppTtreeChapDesc                          uint = 6207
	AppTtreeAngleDesc                         uint = 6210
	AppDvdManualTitle                         uint = 6220
	AppDvdManualText                          uint = 6225
	AppDvdTitlesCount                         uint = 6221
	AppDvdCountCells                          uint = 6222
	AppDvdCountPgc                            uint = 6223
	AppDvdBrokenTitleEntry                    uint = 6224
	AppSingleDriveTitle                       uint = 6226
	AppSingleDriveText                        uint = 6227
	AppSingleDriveAll                         uint = 6228
	AppSingleDriveCaption                     uint = 6229
	AppSiDriveinfo                            uint = 6300
	AppSiProfile                              uint = 6301
	AppSiManufacturer                         uint = 6302
	AppSiProduct                              uint = 6303
	AppSiRevision                             uint = 6304
	AppSiSerial                               uint = 6305
	AppSiFirmware                             uint = 6306
	AppSiFirdate                              uint = 6307
	AppSiBecflags                             uint = 6308
	AppSiHighestAacs                          uint = 6309
	AppSiDiscinfo                             uint = 6320
	AppSiNodisc                               uint = 6321
	AppSiDiscload                             uint = 6322
	AppSiCapacity                             uint = 6323
	AppSiDisctype                             uint = 6324
	AppSiDiscsize                             uint = 6325
	AppSiDiscrate                             uint = 6326
	AppSiDisclayers                           uint = 6327
	AppSiDisccbl                              uint = 6329
	AppSiDisccbl25                            uint = 6330
	AppSiDisccbl27                            uint = 6331
	AppSiDevice                               uint = 6332
)
