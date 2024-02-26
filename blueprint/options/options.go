package options

import (
	"log"
	"os"
	"runtime"
)

var (
	logger = log.Default()
	scaffoldPrefix = ""
)

func init() {
	logger.SetFlags(log.LstdFlags & ^log.Ltime & ^log.Ldate)
}

func DisableLogging() {
	nullDevice := getNullDevice()
	if nullDevice != nil {
		logger.SetOutput(nullDevice)
	}
}

func SetScaffoldPrefix(prefix string) {
	scaffoldPrefix = prefix
}

func GetScaffoldPrefix() *string {
	return &scaffoldPrefix
}

func GetLogger() *log.Logger {
	return logger
}

func getNullDevice() *os.File {
	var nullFileName string
	switch runtime.GOOS {
	case "linux", "darwin", "freebsd", "openbsd", "netbsd":
		nullFileName = "/dev/null"
	case "windows":
		nullFileName = "nul"
	default:
		return nil
	}

	f, _ := os.OpenFile(nullFileName, os.O_RDWR, 0)
	return f
}

