package chast

import (
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

type Chast struct {
	mode    string
	storeID string
	key     string

	apiClient *api.APIClient
}

const (
	testURL  = "https://u2-demo-ext.mono.st4g3.com/api"
	stageURL = "https://u2-ext.mono.st4g3.com/api"
	prodURL  = "https://u2.monobank.com.ua/api"
)

type Opts struct {
	// Mode - режим роботи, може бути test, stage, prod
	Mode    string
	StoreID string
	Key     string
	Client  *http.Client
}

func NewChast(opts Opts) *Chast {
	var url string
	switch opts.Mode {
	case "test":
		url = testURL
	case "stage":
		url = stageURL
	case "prod":
		url = prodURL
	default:
		url = testURL
	}
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &Chast{
		mode:    opts.Mode,
		storeID: opts.StoreID,
		key:     opts.Key,
		apiClient: &api.APIClient{
			BaseURL:         url,
			Client:          opts.Client,
			ErrorParserFunc: errorParser,
		},
	}
}
