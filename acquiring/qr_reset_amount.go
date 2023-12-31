package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// QrResetAmountResponse запит на скидання суми
type QrResetAmountRequest struct {
	QrID string `json:"qrId"`
}

// QrResetAmountResponse відповідь на запит на скидання суми
type QrResetAmountResponse struct{}

// QrResetAmount запит на скидання суми
func (a *Acquiring) QrResetAmount(ctx context.Context, req QrResetAmountRequest) (*QrResetAmountResponse, error) {
	var resp QrResetAmountResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/qr/reset-amount",
		Headers: map[string]string{
			"X-Token": a.token,
		},
		Response: &resp,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
