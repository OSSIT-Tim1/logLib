package loglib

import (
	"github.com/sirupsen/logrus"
)

func GenerateErrorLog(id, srcType, srcNname, ip, message string) {

	event := Event{
		Id:         id,
		SourceType: srcType,
		SourceName: srcNname,
		Ip:         ip,
		EventType:  ERROR,
		Message:    message,
	}
	WriteLog(&event)

}

func GenerateWarningLog(id, srcType, srcNname, ip, message string) {
	event := Event{
		Id:         id,
		SourceType: srcType,
		SourceName: srcNname,
		Ip:         ip,
		EventType:  WARNING,
		Message:    message,
	}
	WriteLog(&event)
}
func GenerateSuccessLog(id, srcType, srcNname, ip, message string) {
	event := Event{
		Id:         id,
		SourceType: srcType,
		SourceName: srcNname,
		Ip:         ip,
		EventType:  SUCCESS,
		Message:    message,
	}
	WriteLog(&event)
}
func GenerateInfoLog(id, srcType, srcNname, ip, message string) {
	event := Event{
		Id:         id,
		SourceType: srcType,
		SourceName: srcNname,
		Ip:         ip,
		EventType:  INFO,
		Message:    message,
	}
	WriteLog(&event)
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}
	// standardLogger.standardFields{Event}
	standardLogger.Formatter = &logrus.TextFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	// invalidArgMessage      = Event{id, "Invalid arg: %s"}
	invalidArgMessage      = "Invalid arg: %s"
	invalidArgValueMessage = "Invalid value for argument: %s: %v"
	missingArgMessage      = "Missing arg: %s"
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage, argumentName)
}

// func main() {
// 	//id, service, username, ip, message string
// 	GenerateInfoLog("1", data.USER_ACCOUNT, "mareerste", "122.22.34.4:5555", "Invalid username")
// }
