package logs

import (
	"fmt"
	"github.com/kataras/golog"
)

// WriteLogs is a function that writes any log to the application file.
// All you have to do is to fill in the code, which corresponds to ctx.StatusCode(XXX),
// and the content, in string format, which corresponds to the text added to the inserted
// log line. The format of the logs used is the following:
// date format: dd/mm/yyyy
// [INFO/WARN/ERR/DBUG] 02/01/2006 15:04:05 Content for log
func WriteLogs(level string, content string) {
	InitLogs()
	fmt.Println("level: " + level)
	switch level {
	case "info":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	case "warn":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	case "error":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	case "debug":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	case "fatal":
		golog.SetLevel("info")
		golog.Infof(content)
		break
	}
}
