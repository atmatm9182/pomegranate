package options

import (
	"io"
	"log"
	"os"
	"runtime"
)

const DefaultScaffoldPrefix = "./"

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

type ScaffoldingOptions struct {
	EnableLogging  bool
	EnableCaching  bool
	ScaffoldPrefix string
}

func DefaultScaffolding() ScaffoldingOptions {
	return ScaffoldingOptions{
		EnableLogging:  true,
		EnableCaching:  false,
		ScaffoldPrefix: "",
	}
}

func (s *ScaffoldingOptions) GetLogger() *log.Logger {
	var outputDevice io.Writer
	if s.EnableLogging {
		outputDevice = os.Stdout
	} else {
		outputDevice = getNullDevice()
		if outputDevice == nil {
			outputDevice = os.Stdout
		}
	}

	flags := log.LstdFlags & ^log.Ldate & ^log.Ltime
	return log.New(outputDevice, "", flags)
}
