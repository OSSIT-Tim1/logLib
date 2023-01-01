package loglib

import (
	"fmt"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	// log "github.com/sirupsen/logrus"
)

// datum i vreme
// izvor dogadjaja - program, komponenta, korisnicki nalog
// tip dogadjaja - error, warning, success
// id dogadjaja
// poruka - detaljniji opis za logovanje

func rotateLog() {
	rl, err := rotatelogs.New(
		"logs/logfile.%Y%m%d%H%M", // for daily rotation we would set "logfile.%Y%m%d"
		rotatelogs.WithLinkName("logfile"),
		rotatelogs.WithMaxAge(24*time.Hour),      // 7*24*time.Hour
		rotatelogs.WithRotationTime(time.Minute), //24*time.Hour
	)
	if err != nil {
		logrus.Fatal(err)
	}

	// Use the rotatelogs object as the io.Writer for the log package.
	logrus.SetOutput(rl)
}

func WriteLog(event *Event) {
	standardFields := logrus.Fields{
		"Log id":      event.Id,
		"Source":      event.SourceType,
		"Source name": event.SourceName,
		"Sender ip":   event.Ip,
		"Event type":  event.EventType,
	}
	rotateLog()
	// log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info(invalidArgument)
	fmt.Println("LOGLIB IZ HELPERA")
	switch event.EventType {
	case ERROR:
		logrus.WithFields(standardFields).Error(event.Message)
	case WARNING:
		logrus.WithFields(standardFields).Warning(event.Message)
	case SUCCESS:
		logrus.WithFields(standardFields).Info(event.Message)
	case INFO:
		logrus.WithFields(standardFields).Info(event.Message)
	default:
		logrus.WithFields(standardFields).Info("EVENT TYPE ERROR")
	}
}

func SaveLog(msg string, logs []string) ([]string, error) {

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
	standardFields := logrus.Fields{
		"Log id":      "1",
		"Source":      "source test",
		"Source name": "source name test",
		"Sender ip":   "source ip test",
	}
	// rotateLog()
	logger.WithFields(standardFields).Info(msg)

	if len(sliceWriter.logs) > 2 {
		fmt.Println("Veci je od 2")
		res, err := flushLogs(logger, sliceWriter.logs)
		if err != nil {
			return sliceWriter.logs, err
		}
		sliceWriter.logs = res
	} else {
		fmt.Println("Nije veci od 2")
	}
	return sliceWriter.logs, nil
}

func flushLogs(logger *logrus.Logger, logs []string) ([]string, error) {

	file, err := os.OpenFile("/data/log/logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Write the logs to the file
	fmt.Println("Pocetak pisanja sa bufferom")
	for br, log := range logs {
		fmt.Println("U for petlji: ", br, log)
		_, err = file.Write([]byte(log))
		if err != nil {
			fmt.Println(err)
			return logs, err
		}
	}

	return []string{}, nil
}
