package logs

import (
	"fmt"
	"github.com/kataras/golog"
	"os"
	"time"
)

func InitLogs() {
	// initialize a new logger
	golog.SetTimeFormat(time.ANSIC)
	//golog.SetLevel("debug")

	golog.ErrorText("|ERROR|", 31)
	golog.WarnText("|WARN|", 32)
	golog.InfoText("|INFO|", 34)
	golog.DebugText("|DEBUG|", 33)

	// TODO : make correct defer func | handle errors !
	/*debugFile, OpenDebugLogFileErr := os.OpenFile("logs/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if OpenDebugLogFileErr != nil {
		golog.Fatalf("Error during opening debug log file: %+v", OpenDebugLogFileErr.Error())
		panic(OpenDebugLogFileErr.Error())
	}
	defer debugFile.Close()

	infoFile, OpenInfoLogFileError := os.OpenFile("logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if OpenInfoLogFileError != nil {
		golog.Fatalf("Error during opening info log file: %+v", OpenInfoLogFileError.Error())
		panic(OpenInfoLogFileError.Error())
	}
	defer infoFile.Close()*/

	// set the outputs per log level.
	/*golog.SetLevelOutput("debug", debugFile)
	golog.SetLevelOutput("info", infoFile)*/

	fmt.Print("\nInitLog -> OK\n")
}

// NewLogFile make new log file
func NewLogFile() *os.File {
	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		golog.Fatal(err)
		WriteLogs("error", err.Error())
		panic(err)
	}

	return f
}

// todayFilename get a filename based on the date, just for the sugar.
func todayFilename() string {
	//today := time.Now().Format("Jan 02 2006")
	today := time.Now().Format("20060102")
	return "CX-20-api-" + today + ".log"
}
