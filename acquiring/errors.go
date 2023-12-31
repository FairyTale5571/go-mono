package acquiring

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/fairytale5571/go-mono/internal/api"
)

type ErrorBody struct {
	ErrCode string `json:"errCode"`
	ErrText string `json:"errText"`
}

func (e ErrorBody) Error() string {
	return fmt.Sprintf("Code %s: %s", e.ErrCode, e.ErrText)
}

var (
	// ErrFailedToParsePEMBlock failed to parse PEM block containing the publicmono key
	ErrFailedToParsePEMBlock = errors.New("failed to parse PEM block containing the publicmono key")

	// ErrFailedToParsePublicKey failed to parse publicmono key
	ErrFailedToParsePublicKey = errors.New("failed to parse publicmono key")
)

var errorParser api.ErrorParserFunc = func(body []byte) error {
	var errBody ErrorBody
	if err := json.Unmarshal(body, &errBody); err != nil {
		return err
	}
	return errors.New(errBody.Error())
}
