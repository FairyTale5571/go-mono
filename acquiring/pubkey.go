package acquiring

import (
	"context"
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// PubKeyResponse відповідь на запит публічного ключа
type PubKeyResponse struct {
	Key string `json:"key"`
}

// PubKey отримання публічного ключа
func (a *Acquiring) PubKey(ctx context.Context) (*PubKeyResponse, error) {
	var response PubKeyResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method:   http.MethodGet,
		Path:     "/pubkey",
		Response: &response,
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
