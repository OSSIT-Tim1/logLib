package loglib

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
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
		log.Fatal(err)
	}

	// Use the rotatelogs object as the io.Writer for the log package.
	log.SetOutput(rl)
}

func WriteLog(event *Event) {
	standardFields := log.Fields{
		"Log id":      event.Id,
		"Source":      event.SourceType,
		"Source name": event.SourceName,
		"Sender ip":   event.Ip,
		"Event type":  event.EventType,
	}
	rotateLog()
	// log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info(invalidArgument)
	switch event.EventType {
	case ERROR:
		log.WithFields(standardFields).Error(event.Message)
	case WARNING:
		log.WithFields(standardFields).Warning(event.Message)
	case SUCCESS:
		log.WithFields(standardFields).Info(event.Message)
	case INFO:
		log.WithFields(standardFields).Info(event.Message)
	default:
		log.WithFields(standardFields).Info("EVENT TYPE ERROR")
	}
}
