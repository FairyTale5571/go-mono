package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type APIClient struct {
	BaseURL string
	Client  *http.Client
}

// request конфігурація запиту
type Request struct {
	Method   string
	Path     string
	Params   map[string]string
	Headers  map[string]string
	Body     any
	Response any
}

// ErrorBody Тіло помилки
type ErrorBody struct {
	ErrCode string `json:"errCode"`
	ErrText string `json:"errText"`
}

// prepareRequestBody Підготовлює тіло запиту
func (a *APIClient) prepareRequestBody(req Request) (io.ReadCloser, error) {
	if req.Body == nil {
		return http.NoBody, nil
	}
	bodyBytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(bytes.NewReader(bodyBytes)), nil
}

// isNotOkStatus Перевіряє чи не є статус помилковим (не 200-299)
func isNotOkStatus(status int) bool {
	return status < http.StatusOK || status > http.StatusMultipleChoices
}

// sendRequest Відправляє запит до API
func (a *APIClient) SendRequest(ctx context.Context, r Request) error {
	reqURL, err := url.Parse(fmt.Sprintf("%s%s", a.BaseURL, r.Path))
	if err != nil {
		return err
	}
	reqURL.Path = r.Path
	values := url.Values{}
	for k, v := range r.Params {
		values.Add(k, v)
	}
	reqURL.RawQuery = values.Encode()

	body, err := a.prepareRequestBody(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, r.Method, reqURL.String(), body)
	if err != nil {
		return err
	}

	for k, v := range r.Headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := a.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if isNotOkStatus(resp.StatusCode) {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var errorBody ErrorBody
		if err := json.Unmarshal(bodyBytes, &errorBody); err != nil {
			return err
		}
		return fmt.Errorf("bad status code: %d, errCode: %s, errText: %s", resp.StatusCode, errorBody.ErrCode, errorBody.ErrText)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, &r.Response)
}
