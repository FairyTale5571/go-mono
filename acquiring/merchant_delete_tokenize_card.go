package acquiring

import (
	"context"
	"github.com/fairytale5571/go-mono/internal/api"
	"net/http"
)

type MerchantDeleteTokenizeCardResponse struct{}

// MerchantDeleteTokenizeCard видалення токенізованої картки
func (a *Acquiring) MerchantDeleteTokenizeCard(ctx context.Context, cardToken string) error {
	return a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodDelete,
		Path:   "/wallet/card",
		Headers: map[string]string{
			"X-Token": a.token,
		},
		Body: map[string]interface{}{
			"cardToken": cardToken,
		},
	})
}
