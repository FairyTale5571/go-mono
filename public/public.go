package public

import (
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// Public - Клієнт для публічного API
type Public struct {
	token     string
	apiClient *api.APIClient
}

const (
	publicApiUrl = "https://api.monobank.ua"
)

type Opts struct {
	Token  string
	Client *http.Client
}

// NewPublic - Створення нового клієнта для публічного API
func NewPublic(opts Opts) *Public {
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &Public{
		token: opts.Token,
		apiClient: &api.APIClient{
			Client:  opts.Client,
			BaseURL: publicApiUrl,
		},
	}
}
