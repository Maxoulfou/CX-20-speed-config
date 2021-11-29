package logs

import (
	"cx-20-api/entity"
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
func WriteLogs(level string, content string, console bool) {
	// load yml conf
	// envConfig := configuration.GetEnv()
	var YamlEnv entity.YmlConfig
	YamlEnv.GetConfig()

	switch level {
	case "error":
		//ErrorFile := NewLogFile("ERROR")
		ErrorFile := &os.File{}
		if YamlEnv.Log == "separated" {
			ErrorFile = NewLogFile("ERROR")
		} else if YamlEnv.Log == "all-in-one" {
			ErrorFile = NewLogFile("log")
		}

		defer func(ErrorFile *os.File) {
			DeferNewLogFileErr := ErrorFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(ErrorFile)
		golog.SetLevelOutput("error", ErrorFile)
		break
	case "warn":
		//WarnFile := NewLogFile("WARN")
		WarnFile := &os.File{}
		if YamlEnv.Log == "separated" {
			WarnFile = NewLogFile("WARN")
		} else if YamlEnv.Log == "all-in-one" {
			WarnFile = NewLogFile("log")
		}

		defer func(WarnFile *os.File) {
			DeferNewLogFileErr := WarnFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(WarnFile)
		golog.SetLevelOutput("warn", WarnFile)
		break
	case "info":
		//InfoFile := NewLogFile("INFO")
		InfoFile := &os.File{}
		if YamlEnv.Log == "separated" {
			InfoFile = NewLogFile("INFO")
		} else if YamlEnv.Log == "all-in-one" {
			InfoFile = NewLogFile("log")
		}

		defer func(InfoFile *os.File) {
			DeferNewLogFileErr := InfoFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(InfoFile)
		golog.SetLevelOutput("info", InfoFile)
		break
	case "debug":
		//DebugFile := NewLogFile("DEBUG")
		DebugFile := &os.File{}
		if YamlEnv.Log == "separated" {
			DebugFile = NewLogFile("DEBUG")
		} else if YamlEnv.Log == "all-in-one" {
			DebugFile = NewLogFile("log")
		}

		defer func(DebugFile *os.File) {
			DeferNewLogFileErr := DebugFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(DebugFile)
		golog.SetLevelOutput("debug", DebugFile)
		break
	case "fatal":
		//FatalFile := NewLogFile("FATAL")
		FatalFile := &os.File{}
		if YamlEnv.Log == "separated" {
			FatalFile = NewLogFile("FATAL")
		} else if YamlEnv.Log == "all-in-one" {
			FatalFile = NewLogFile("log")
		}

		defer func(FatalFile *os.File) {
			DeferNewLogFileErr := FatalFile.Close()
			if DeferNewLogFileErr != nil {
				log.Fatalf("New log file : " + DeferNewLogFileErr.Error())
			}
		}(FatalFile)
		golog.SetLevelOutput("fatal", FatalFile)
		break
	default:
		golog.Fatal("Any level of log are passed in arguments !")
		os.Exit(1)
		break
	}

	// golog.SetOutput(io.MultiWriter(f, os.Stdout))

	golog.SetTimeFormat(time.ANSIC)

	golog.ErrorText("[ERROR]", 31)
	golog.WarnText("[WARN]", 32)
	golog.InfoText("[INFO]", 34)
	golog.DebugText("[DEBUG]", 33)
	// Can't modify fatal log level displaying text

	// DEBUG : fmt.Println("level: " + level)
	switch level {
	case "info":
		golog.SetLevel("info")
		golog.Infof(content)
		if console {
			fmt.Printf(level + " : " + content)
		}
		break
	case "warn":
		golog.SetLevel("warn")
		golog.Warnf(content)
		if console {
			fmt.Printf(level + " : " + content)
		}
		break
	case "error":
		golog.SetLevel("error")
		golog.Errorf(content)
		if console {
			fmt.Printf(level + " : " + content)
		}
		break
	case "debug":
		golog.SetLevel("debug")
		golog.Debugf(content)
		if console {
			fmt.Printf(level + " : " + content)
		}
		break
	case "fatal":
		golog.SetLevel("fatal")
		golog.Fatalf(content)
		if console {
			fmt.Printf(level + " : " + content)
		}
		break
	}
}
