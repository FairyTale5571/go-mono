package personal

import (
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
	openapi "github.com/FairyTale5571/go-mono/open-api"
)

// Personal - Клієнт для публічного API
type Personal struct {
	token     string
	apiClient *api.APIClient
}

const (
	personalAPIURL = "https://api.monobank.ua"
)

type Opts struct {
	Token  string
	Client *http.Client
}

// NewPersonalClient - Створення нового клієнта для публічного API
func NewPersonalClient(opts Opts) *Personal {
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &Personal{
		token: opts.Token,
		apiClient: &api.APIClient{
			Client:          opts.Client,
			BaseURL:         personalAPIURL,
			ErrorParserFunc: openapi.ErrorParser,
		},
	}
}
