package loglib

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// const maxMemory = 1024 * 1024 // boundary of 1MB for logfile

/*
The saveLog function takes two parameters (forwarded logs and Event object with parameters)
When new log was created, it is inserted into existing logs. If logs are above agreed boundary,
logs would be written down to local logfile depending on the time when log was created.
If logs are under boundary, they are going to be returned with new inserted log.
*/
func saveLog(logs []string, event *Event) ([]string, error) {

	logger := logrus.New()
	sliceWriter := &SliceWriter{
		logs: logs,
	}
	logger.Out = sliceWriter
	// Set the log level to debug.
	logger.SetLevel(logrus.DebugLevel)
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC822,
		FullTimestamp:   true,
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return sliceWriter.logs, err
	}
	standardFields := logrus.Fields{
		"Log id":      id,
		"Source name": event.SourceName,
		"Source type": event.SourceType,
		"Sender ip":   event.Ip,
	}

	switch event.EventType {
	case ERROR:
		logger.WithFields(standardFields).Error(event.Message)
	case WARNING:
		logger.WithFields(standardFields).Warning(event.Message)
	case SUCCESS:
		logger.WithFields(standardFields).Info(event.Message)
	case INFO:
		logger.WithFields(standardFields).Info(event.Message)
	default:
		logger.WithFields(standardFields).Info(event.Message)
	}

	if len(sliceWriter.logs) > 10 { // should be replaced with maxMemory
		res, err := flushLogs(logger, sliceWriter.logs)
		if err != nil {
			return sliceWriter.logs, err
		}
		sliceWriter.logs = res
	}
	return sliceWriter.logs, nil
}

// Write the logs to the file
func flushLogs(logger *logrus.Logger, logs []string) ([]string, error) {

	newLogs := make([]string, 0, len(logs))
	for _, log := range logs {
		err := rotateLog(log)
		if err != nil {
			newLogs = append(newLogs, log)
		}

	}

	return newLogs, nil
}

/*
Tge rotateLog function is used for creating new and update existing log files
which differ in our case by the minutes of creation (should be the day of creation).
*/
func rotateLog(log string) error {
	date, err := extractDateFromLog(log)
	if err != nil {
		return err
	}

	rl, err := rotatelogs.New(
		fmt.Sprintf("/data/log/logfile.%s", date), //for daily rotation we would set "logfile.%Y%m%d"
		rotatelogs.WithMaxAge(24*time.Hour),       // 7*24*time.Hour
		rotatelogs.WithRotationTime(time.Minute),  //24*time.Hour
	)
	if err != nil {
		return err
	}

	_, err = rl.Write([]byte(log)) //write in file
	if err != nil {
		return err
	}

	return nil
}

func extractDateFromLog(log string) (string, error) {
	// Extract the date and time from the recived log (string)
	re := regexp.MustCompile(`time="(.+)" level`)
	dateTimeString := re.FindStringSubmatch(log)[1]

	// Parse the date and time string using the time.RFC822 layout
	t, err := time.Parse(time.RFC822, dateTimeString)
	if err != nil {
		return "", err
	}
	formatted := t.Format("200601021504")

	// Print the formatted date and time
	return formatted, nil
}
