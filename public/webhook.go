package public

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

type SetWebhookRequest struct {
	WebHookURL string `json:"webHookUrl"`
}

func (p *Public) SetWebhook(ctx context.Context, req *SetWebhookRequest) error {
	return p.apiClient.SendRequest(ctx,api.Request{
		Method: http.MethodPost,
		Path:   "/personal/webhook",
		Body:   req,
		Headers: map[string]string{
			"X-Token": p.token,
		},
	})
}
