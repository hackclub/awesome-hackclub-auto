package logging

import (
	"fmt"
	"os"
	"strings"

	"github.com/Matt-Gleich/logoru"
	"github.com/honeybadger-io/honeybadger-go"
)

func Log(v interface{}, level string, localOnly bool) {
	// Only log to Honeybadger if an API key is set and it's an error
	if os.Getenv("HONEYBADGER_API_KEY") != "" && !localOnly && (level == "error" || level == "critical" || level == "warning") {
		logoru.Info("Logging an error to Honeybadger...")
		_, err := honeybadger.Notify(fmt.Sprintf("%s: %v", strings.ToUpper(level), v))
		if err != nil {
			logoru.Error("Ironically enough, something went wrong while trying to log an error :(")
		}
	}

	switch level {
	case "debug":
		logoru.Debug(v)
	case "info":
		logoru.Info(v)
	case "warning":
		logoru.Warning(v)
	case "error":
		logoru.Error(v)
	case "critical":
		logoru.Critical(v)
	case "success":
		logoru.Success(v)
	}
}