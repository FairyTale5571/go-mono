package corp

import (
	"context"
	"net/http"
	"time"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// AuthRegistrationRequest тіло запиту на реєстрацію
type AuthRegistrationRequest struct {
	Pubkey        string `json:"pubkey"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ContactPerson string `json:"contactPerson"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Logo          string `json:"logo"`
}

// AuthRegistrationResponse тіло відповіді на реєстрацію
type AuthRegistrationResponse struct {
	Status string `json:"status"`
}

// authRegistration - Створення заявки на авторизацію компанії банком
func (c *Corp) authRegistration(ctx context.Context, req AuthRegistrationRequest) (*AuthRegistrationResponse, error) {
	var resp AuthRegistrationResponse
	err := c.apiClient.SendRequest(ctx, api.Request{
		Method:   http.MethodPost,
		Path:     "/auth/register",
		Response: &resp,
		Headers: map[string]string{
			"X-Sign": c.token,
			"X-Time": time.Now().UTC().String(),
		},
		Body: req,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
