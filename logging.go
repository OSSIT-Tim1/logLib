package loglib

var WrongCredentialsMsg = "User with username %s attempted to submit incorrect credentials"
var ExistingCredentialsMsg = "User with username %s trying to sign up with existing credentials"
var MultipleAttemptSignInMsg = "User with username %s trying to sign in more then %d times with wrong credentials"
var ItemCreated = "A new %s has been created"
var ItemUpdated = "%s has been updated"
var ItemDeleted = "%s has been deleted"
var EmailSent = "A new mail has been created for %s and sent to his/her address"
var SuccessfulSignIn = "User with username %s has successfully logged in"
var InvalidJSON = "Invalid JSON format sent"
var DataDoesntExist = "Required data doesn't exist"

func GenerateLog(msg string, logs []string, logType, sourceType, sourceName, ip string) ([]string, error) {
	event := Event{
		SourceType: sourceType,
		SourceName: sourceName,
		Ip:         ip,
		EventType:  logType,
		Message:    msg,
	}

	res, err := saveLog(logs, &event)
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
	res, err := saveLog(logs, &event)
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
	res, err := saveLog(logs, &event)
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
	res, err := saveLog(logs, &event)
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
	res, err := saveLog(logs, &event)
	if err != nil {
		return res, err
	}
	return res, nil
}
