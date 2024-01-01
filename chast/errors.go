package chast

import (
	"encoding/json"
	"errors"

	"github.com/FairyTale5571/go-mono/internal/api"
)

type ErrorBody struct {
	ErrorDescription string `json:"errorDescription"`
}

func (e ErrorBody) Error() string {
	return e.ErrorDescription
}

var errorParser api.ErrorParserFunc = func(body []byte) error {
	var errBody ErrorBody
	if err := json.Unmarshal(body, &errBody); err != nil {
		return err
	}
	return errors.New(errBody.Error())
}
