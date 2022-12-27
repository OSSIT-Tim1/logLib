package loglib

const (
	ERROR   = "Error"
	WARNING = "Warning"
	SUCCESS = "Success"
	INFO    = "Info"

	SERVICE      = "Service"
	USER_ACCOUNT = "Username"
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
