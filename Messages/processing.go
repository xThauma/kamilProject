package Messages

import (
	"net/http"
	"strconv"
)

func FormToMessage(r *http.Request) (Message, []string) {
	var message Message
	var errStr, magic_numberStr string
	var errs []string
	var err error

	message.Email, errStr = processFormField(r, "email")
	errs = appendError(errs, errStr)
	message.Title, errStr = processFormField(r, "title")
	errs = appendError(errs, errStr)
	message.Content, errStr = processFormField(r, "content")
	errs = appendError(errs, errStr)

	magic_numberStr, errStr = processFormField(r, "magic_number")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		message.Magic_number, err = strconv.Atoi(magic_numberStr)
		if err != nil {
			errs = append(errs, "Parameter 'magic_number' not an integer")
		}
	}
	return message, errs
}

func appendError(errs []string, errStr string) ([]string) {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}

