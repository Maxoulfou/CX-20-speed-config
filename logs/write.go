package logs

import (
	"fmt"
	"github.com/kataras/golog"
	"log"
	"os"
	"time"
)

// WriteLogs is a function that writes any log to the application file.
// All you have to do is to fill in the code, which corresponds to ctx.StatusCode(XXX),
// and the content, in string format, which corresponds to the text added to the inserted
// log line. The format of the logs used is the following:
// date format: dd/mm/yyyy
// [INFO/WARN/ERR/DBUG] 02/01/2006 15:04:05 Content for log
func WriteLogs(level string, content string) {
	// TODO : make correct defer func | handle errors !

	switch level {
	case "error":
		ErrorFile := NewLogFile("ERROR")
		defer func(ErrorFile *os.File) {
			DeferNewLogFileErr := ErrorFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(ErrorFile)
		golog.SetLevelOutput("error", ErrorFile)
		break
	case "warn":
		WarnFile := NewLogFile("WARN")
		defer func(WarnFile *os.File) {
			DeferNewLogFileErr := WarnFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(WarnFile)
		golog.SetLevelOutput("warn", WarnFile)
		break
	case "info":
		InfoFile := NewLogFile("INFO")
		defer func(InfoFile *os.File) {
			DeferNewLogFileErr := InfoFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(InfoFile)
		golog.SetLevelOutput("info", InfoFile)
		break
	case "debug":
		DebugFile := NewLogFile("DEBUG")
		defer func(DebugFile *os.File) {
			DeferNewLogFileErr := DebugFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(DebugFile)
		golog.SetLevelOutput("debug", DebugFile)
		break
	}

	// golog.SetOutput(io.MultiWriter(f, os.Stdout))

	golog.SetTimeFormat(time.ANSIC)

	golog.ErrorText("[ERROR]", 31)
	golog.WarnText("[WARN]", 32)
	golog.InfoText("[INFO]", 34)
	golog.DebugText("[DEBUG]", 33)

	fmt.Println("level: " + level)
	switch level {
	case "info":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	case "warn":
		golog.SetLevel("warn")
		golog.Warnf(content)
		break
	case "error":
		golog.SetLevel("error")
		golog.Errorf(content)
		break
	case "debug":
		golog.SetLevel("debug")
		golog.Debugf(content)
		break
	case "fatal":
		golog.SetLevel("fatal")
		golog.Fatalf(content)
		break
	}
}
