package openapi

import (
	"encoding/json"
	"errors"

	"github.com/FairyTale5571/go-mono/internal/api"
)

type ErrorBody struct {
	Message string `json:"message"`
}

func (e ErrorBody) Error() string {
	return e.Message
}

var ErrorParser api.ErrorParserFunc = func(body []byte) error {
	var errBody ErrorBody
	if err := json.Unmarshal(body, &errBody); err != nil {
		return err
	}
	return errors.New(errBody.Error())
}
