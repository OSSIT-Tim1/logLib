package loglib

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
