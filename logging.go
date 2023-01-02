package loglib

var WrongCredentialsMsg = "User with username {} trying to sign in with wrong credentials"
var MultipleAttemptSignInMsg = "User with username {} trying to sign in more then {} times with wrong credentials"
var SuccessfulSignIn = "User with username {} has successfully logged in"

func GenerateLog(msg string, logs []string, logType string, sourceType, sourceName, ip string) ([]string, error) {
	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  logType,
		Message:    msg,
	}

	res, err := saveLog(logs,&event)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GenerateErrorLog(msg string, logs []string, sourceType, sourceName, ip string) ([]string, error) {

	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  ERROR,
		Message:    msg,
	}
	res, err := saveLog(logs,&event)
	if err != nil {
		return res, err
	}
	return res, nil

}

func GenerateWarningLog(msg string, logs []string, sourceType, sourceName, ip string) ([]string, error) {

	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  WARNING,
		Message:    msg,
	}
	res, err := saveLog(logs,&event)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GenerateSuccessLog(msg string, logs []string, sourceType, sourceName, ip string) ([]string, error) {

	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  SUCCESS,
		Message:    msg,
	}
	res, err := saveLog(logs,&event)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GenerateInfoLog(msg string, logs []string, sourceType, sourceName, ip string) ([]string, error) {

	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  INFO,
		Message:    msg,
	}
	res, err := saveLog(logs,&event)
	if err != nil {
		return res, err
	}
	return res, nil
}
