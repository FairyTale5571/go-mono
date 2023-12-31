package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// InvoiceRemoveRequest тіло запиту на скасування рахунку
type InvoiceRemoveRequest struct {
	InvoiceID string `json:"invoiceId"`
}

// RemoveInvoice Скасування рахунку
func (a *Acquiring) RemoveInvoice(ctx context.Context, req *InvoiceRemoveRequest) error {
	return a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/remove",
		Headers: map[string]string{
			"X-Token": a.token,
		},
		Body: req,
	})
}
