package expirenza

import (
	"encoding/json"
	"errors"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// ErrInvalidResponseBody is returned when the response body is not valid.
var ErrInvalidResponseBody = errors.New("invalid response body")

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorBody struct {
	ErrCode     interface{}   `json:"errCode"`
	ErrText     string        `json:"errText"`
	ErrTitle    string        `json:"errTitle"`
	ErrType     string        `json:"errType"`
	FieldErrors []*FieldError `json:"fieldErrors"`
}

func (e *ErrorBody) Error() string {
	return e.ErrText
}

var errorParser api.ErrorParserFunc = func(body []byte) error {
	var errBody ErrorBody
	if err := json.Unmarshal(body, &errBody); err != nil {
		return err
	}
	return errors.New(errBody.Error())
}
