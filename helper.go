package loglib

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// datum i vreme
// izvor dogadjaja - program, komponenta, korisnicki nalog
// tip dogadjaja - error, warning, success
// id dogadjaja
// poruka - detaljniji opis za logovanje

func extractDateFromLog(log string) (string, error) {
	// Extract the date and time string from the given string
	re := regexp.MustCompile(`time="(.+)" level`)
	dateTimeString := re.FindStringSubmatch(log)[1]

	// Parse the date and time string using the time.RFC822 layout
	t, err := time.Parse(time.RFC822, dateTimeString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	formatted := t.Format("200601021504")

	// Print the formatted date and time
	return formatted, nil
}

func rotateLog(log string) error {
	date, err := extractDateFromLog(log)
	if err != nil {
		fmt.Println(err)
	}

	rl, err := rotatelogs.New(
		// "logs/logfile.%Y%m%d%H%M", // for daily rotation we would set "logfile.%Y%m%d"
		// fmt.Sprintf("logs/logfile.%s", date), // for daily rotation we would set "logfile.%Y%m%d"
		fmt.Sprintf("/data/log/logfile.%s", date),
		// rotatelogs.WithLinkName("/data/log/logfile"),
		rotatelogs.WithMaxAge(24*time.Hour),      // 7*24*time.Hour
		rotatelogs.WithRotationTime(time.Minute), //24*time.Hour
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//write in file
	_, err = rl.Write([]byte(log))
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Use the rotatelogs object as the io.Writer for the log package.
	// logrus.SetOutput(rl)
	return nil
}

// func WriteLog(event *Event) {
// 	standardFields := logrus.Fields{
// 		"Log id":      event.Id,
// 		"Source":      event.SourceType,
// 		"Source name": event.SourceName,
// 		"Sender ip":   event.Ip,
// 		"Event type":  event.EventType,
// 	}
// 	rotateLog()
// 	// log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info(invalidArgument)
// 	fmt.Println("LOGLIB IZ HELPERA")
// 	switch event.EventType {
// 	case ERROR:
// 		logrus.WithFields(standardFields).Error(event.Message)
// 	case WARNING:
// 		logrus.WithFields(standardFields).Warning(event.Message)
// 	case SUCCESS:
// 		logrus.WithFields(standardFields).Info(event.Message)
// 	case INFO:
// 		logrus.WithFields(standardFields).Info(event.Message)
// 	default:
// 		logrus.WithFields(standardFields).Info("EVENT TYPE ERROR")
// 	}
// }

const maxMemory = 1024 * 1024 // boundary of 1MB for logfile

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
		fmt.Println(err)
		return sliceWriter.logs, err
	}
	standardFields := logrus.Fields{
		"Log id":      id,
		"Source name": event.SourceName,
		"Source type": event.SourceType,
		"Sender ip":   event.Ip,
	}
	// rotateLog()

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

	if len(sliceWriter.logs) > 2 { // 2 should be replaced with maxMemory
		res, err := flushLogs(logger, sliceWriter.logs)
		if err != nil {
			return sliceWriter.logs, err
		}
		sliceWriter.logs = res
	}
	return sliceWriter.logs, nil
}

func flushLogs(logger *logrus.Logger, logs []string) ([]string, error) {

	// Write the logs to the file
	newLogs := make([]string, 0, len(logs))
	for _, log := range logs {
		err := rotateLog(log)
		if err != nil {
			fmt.Println("Error during rotation log: ", err)
			newLogs = append(newLogs, log)
		}

	}

	return newLogs, nil
}
