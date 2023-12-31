package public

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// SetWebhookRequest - Запит на встановлення URL для надсилання подій по зміні балансу рахунку
type SetWebhookRequest struct {
	WebHookURL string `json:"webHookUrl"`
}

// SetWebhook - Встановлення URL для надсилання подій по зміні балансу рахунку
func (p *Public) SetWebhook(ctx context.Context, req *SetWebhookRequest) error {
	return p.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/personal/webhook",
		Body:   req,
		Headers: map[string]string{
			"X-Token": p.token,
		},
	})
}
