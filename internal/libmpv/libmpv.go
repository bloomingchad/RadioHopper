package libmpv

import "C"

type MpvError int

const (
	ErrGeneric             MpvError = -20
	ErrNotImplemented      MpvError = -19
	ErrUnsupported         MpvError = -18
	ErrUnknownFormat       MpvError = -17
	ErrNothingToPlay       MpvError = -16
	ErrVOInitFailed        MpvError = -15
	ErrAOInitFailed        MpvError = -14
	ErrLoadFailed          MpvError = -13
	ErrCmd                 MpvError = -12
	ErrOnProperty          MpvError = -11
	ErrPropertyUnavailable MpvError = -10
	ErrPropertyFormat      MpvError = -9
	ErrPropertyNotFound    MpvError = -8
	ErrOnOption            MpvError = -7
	ErrOptionFormat        MpvError = -6
	ErrOptionNotFound      MpvError = -5
	ErrInvalidParameter    MpvError = -4
	ErrUninitialized       MpvError = -3
	ErrNoMem               MpvError = -2
	ErrEventQueueFull      MpvError = -1
	ErrSuccess             MpvError = 0
)

type MpvFormat int

const (
	FmtNone      MpvFormat = 0
	FmtString    MpvFormat = 1
	FmtOSDString MpvFormat = 2
	FmtFlag      MpvFormat = 3
	FmtInt64     MpvFormat = 4
	FmtFloat64   MpvFormat = 5
	FmtNode      MpvFormat = 6
	FmtNodeArray MpvFormat = 7
	FmtNodeMap   MpvFormat = 8
	FmtByteArray MpvFormat = 9
)

type MpvEventID int

const (
	IDNone                MpvEventID = 0
	IDShutdown            MpvEventID = 1
	IDLogMessage          MpvEventID = 2
	IDGetPropertyReply    MpvEventID = 3
	IDSetPropertyReply    MpvEventID = 4
	IDCommandReply        MpvEventID = 5
	IDStartFile           MpvEventID = 6
	IDEndFile             MpvEventID = 7
	IDFileLoaded          MpvEventID = 8
	IDClientMessage       MpvEventID = 16
	IDVideoReConfig       MpvEventID = 17
	IDAudioReConfig       MpvEventID = 18
	IDSeek                MpvEventID = 20
	IDPlayBackRestart     MpvEventID = 21
	IDEventPropertyChange MpvEventID = 22
	IDQueueOverFlow       MpvEventID = 24
	IDEventHook           MpvEventID = 25
)

type MpvEndFileReason int

const (
	EfrEOF      MpvEndFileReason = 0
	EfrStop     MpvEndFileReason = 2
	EfrQuit     MpvEndFileReason = 3
	EfrError    MpvEndFileReason = 4
	EfrReDirect MpvEndFileReason = 5
)

type MpvLogLevel int

const (
	LlNone    MpvLogLevel = 0
	LlFatal   MpvLogLevel = 10
	LlError   MpvLogLevel = 20
	LlWarn    MpvLogLevel = 30
	LlInfo    MpvLogLevel = 40
	LlVerbose MpvLogLevel = 50
	LlDebug   MpvLogLevel = 60
	LlTrace   MpvLogLevel = 70
)
