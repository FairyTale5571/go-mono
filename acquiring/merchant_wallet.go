package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

type MerchantWallet struct {
	CardToken string `json:"cardToken"`
	MaskedPan string `json:"maskedPan"`
	Country   string `json:"country"`
}

type MerchantWalletsResponse struct {
	Wallets []*MerchantWallet `json:"wallets"`
}

// MerchantWallets отримання списку токенізованих карток
func (a *Acquiring) MerchantWallets(ctx context.Context, walletID string) (*MerchantWalletsResponse, error) {
	var response MerchantWalletsResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/wallet",
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
