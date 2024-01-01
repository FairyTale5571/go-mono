package acquiring

import (
	"context"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// SubMerchant обʼект субмерчантів
type SubMerchant struct {
	Code   string `json:"code"`
	EDRPOU string `json:"edrpou"`
	IBAN   string `json:"iban"`
}

// SubMerchantResponse відповідь на запит субмерчантів
type SubMerchantResponse struct {
	List []*SubMerchant `json:"list"`
}

// MerchantSubMerchant отримання субмерчантів
func (a *Acquiring) MerchantSubMerchant(ctx context.Context) (*SubMerchantResponse, error) {
	var response SubMerchantResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: "GET",
		Path:   "/submerchant/list",
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
