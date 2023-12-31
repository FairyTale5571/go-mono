package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// MerchantDetailsResponse відповідь на запит деталей мерчанта
type MerchantDetailsResponse struct {
	MerchantID   string `json:"merchantId"`
	MerchantName string `json:"merchantName"`
	EDRPOU       string `json:"edrpou"`
}

// MerchantDetails отримання деталей мерчанта
func (a *Acquiring) MerchantDetails(ctx context.Context) (*MerchantDetailsResponse, error) {
	var response MerchantDetailsResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/details",
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
