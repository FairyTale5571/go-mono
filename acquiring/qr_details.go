package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// QrDetailsResponse Інформація про QR-касу, лише для активованих QR-кас
type QrDetailsResponse struct {
	ShortQrID string `json:"shortQrId"`
	InvoiceID string `json:"invoiceId"`
	Amount    int    `json:"amount"`
	Ccy       int    `json:"ccy"`
}

// QrDetails Інформація про QR-касу, лише для активованих QR-кас
func (a *Acquiring) QrDetails(ctx context.Context, qrID string) (*QrDetailsResponse, error) {
	var response QrDetailsResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/qr/details",
		Params: map[string]string{
			"qrId": qrID,
		},
		Headers: map[string]string{
			"X-Token": a.token,
		},
		Response: &response,
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
