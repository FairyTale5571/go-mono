package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// QrObject QR-каса
type QrObject struct {
	ShortQrID  string `json:"shortQrId"`
	QrID       string `json:"qrId"`
	AmountType string `json:"amountType"`
	PageURL    string `json:"pageUrl"`
}

// QrListResponse список QR-кас
type QrListResponse struct {
	List []*QrObject `json:"list"`
}

// QrList список QR-кас
func (a *Acquiring) QrList(ctx context.Context) (*QrListResponse, error) {
	var response QrListResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/qr/list",
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
