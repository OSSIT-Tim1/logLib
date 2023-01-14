package loglib

const (
	ERROR   = "Error"
	WARNING = "Warning"
	SUCCESS = "Success"
	INFO    = "Info"

	SERVICE      = "Service"
	USER_ACCOUNT = "User"
	COMPONENT    = "Component"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	Id         string
	SourceType string
	SourceName string
	Ip         string
	EventType  string
	Message    string
}

type SliceWriter struct {
	logs []string
}

func (sw *SliceWriter) Write(p []byte) (int, error) {
	sw.logs = append(sw.logs, string(p))
	return len(p), nil
}
