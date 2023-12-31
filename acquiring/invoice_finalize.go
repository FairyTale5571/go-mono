package acquiring

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

// InvoiceFinalizeRequest тіло запиту на завершення рахунку
type InvoiceFinalizeRequest struct {
	InvoiceID string `json:"invoiceId"`
	Amount    int64  `json:"amount"`
}

// InvoiceFinalizeResponse тіло відповіді на завершення рахунку
type InvoiceFinalizeResponse struct {
	Status string `json:"status"`
}

// FinalizeInvoice Завершення рахунку
func (a *Acquiring) FinalizeInvoice(ctx context.Context, req *InvoiceFinalizeRequest) (*InvoiceFinalizeResponse, error) {
	var resp InvoiceFinalizeResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/finalize",
		Body:   req,
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
