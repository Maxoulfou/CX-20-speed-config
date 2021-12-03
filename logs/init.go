package logs

import (
	"errors"
	"github.com/kataras/golog"
	"log"
	"os"
	"time"
)

/*
	Error
	Warn
	Info
	Debug
*/

// CheckLogsDirectory will check if 'logs' directory exist, if not, it will crete them
func CheckLogsDirectory() {
	path := "logs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

// NewLogFile make new log file
func NewLogFile(loglevel string) *os.File {
	CheckLogsDirectory()

	filename := todayFilename()
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile("logs/"+loglevel+"-"+filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		golog.Fatal(err)
		WriteLogs("error", err.Error(), true)

		panic(err)
	}

	return f
}

// todayFilename get a filename based on the date, just for the sugar.
func todayFilename() string {
	//today := time.Now().Format("Jan 02 2006")
	today := time.Now().Format("02_01_2006")

	return "CX-20-api-" + today + ".log"
}
