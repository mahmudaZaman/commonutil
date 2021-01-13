package resterror

import (
	"github.com/mandatorySuicide/ts-common/accessory"
	"net/http"
	"time"
)

/*
DisplayMessage contains message title,description and code.
This message will de displayed to end user, please choose the words carefully.
code will be used for multilingual purpose.
*/
type DisplayMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

/*
RestError represents rest api error message to client.
Only service should return this which will later be consumed by controller.
utility, repository or any other function should return normal error.
*/
type RestError struct {
	DisplayMessage *DisplayMessage `json:"displayMsg"`
	HTTPStatus     int             `json:"httpStatus"`
	Error          error           `json:"error"`
	Timestamp      string          `json:"timestamp"`
}

// NewBadRequestError function creates a rest error with status 400, intended to be used to notify client in case of API-UI contract breach.
func NewBadRequestError(dmTitle, dmDescription, dmCode string, err error) *RestError {
	return &RestError{
		DisplayMessage: &DisplayMessage{
			Title:       dmTitle,
			Description: dmDescription,
			Code:        dmCode,
		},
		HTTPStatus: http.StatusBadRequest,
		Error:      err,
		Timestamp:  accessory.Func.APIDateFormat(time.Now().UTC()),
	}
}

/*
NewInternalServerError function creates a rest error with status 500,
intended to be used when server encounter unknown/handleable/ambiguous error.
*/
func NewInternalServerError(dmTitle, dmDescription, dmCode string, err error) *RestError {
	return &RestError{
		DisplayMessage: &DisplayMessage{
			Title:       dmTitle,
			Description: dmDescription,
			Code:        dmCode,
		},
		HTTPStatus: http.StatusInternalServerError,
		Error:      err,
		Timestamp:  accessory.Func.APIDateFormat(time.Now().UTC()),
	}
}

/*
NewUnAuthorizedError function creates a rest error with status 401,
intended to be mostly with unauthenticated users trying to access an endpoint that requires authentication.
*/
func NewUnAuthorizedError(dmTitle, dmDescription, dmCode string, err error) *RestError {
	return &RestError{
		DisplayMessage: &DisplayMessage{
			Title:       dmTitle,
			Description: dmDescription,
			Code:        dmCode,
		},
		HTTPStatus: http.StatusUnauthorized,
		Error:      err,
		Timestamp:  accessory.Func.APIDateFormat(time.Now().UTC()),
	}
}

/*
NewForbiddenError function creates a rest error with status 403,
intended to be mostly with authenticated users trying to access an endpoint that requires additional privilege.
*/
func NewForbiddenError(dmTitle, dmDescription, dmCode string, err error) *RestError {
	return &RestError{
		DisplayMessage: &DisplayMessage{
			Title:       dmTitle,
			Description: dmDescription,
			Code:        dmCode,
		},
		HTTPStatus: http.StatusForbidden,
		Error:      err,
		Timestamp:  accessory.Func.APIDateFormat(time.Now().UTC()),
	}
}
