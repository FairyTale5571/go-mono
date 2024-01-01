package acquiring

import (
	"context"
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// Fiscal обʼєкт фіскалізації
type Fiscal struct {
	ID                  string `json:"id"`
	Type                string `json:"type"`
	Status              string `json:"status"`
	StatusDescription   string `json:"statusDescription"`
	TaxURL              string `json:"taxUrl"`
	File                string `json:"file"`
	FiscalizationSource string `json:"fiscalizationSource"`
}

// FiscalResponse відповідь на запит фіскалізації
type FiscalResponse struct {
	Checks []*Fiscal `json:"checks"`
}

// InvoiceFiscal фіскалізація чека
func (a *Acquiring) InvoiceFiscal(ctx context.Context, invoiceID string) (*FiscalResponse, error) {
	var response FiscalResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/fiscal-checks",
		Headers: map[string]string{
			"X-Token": a.token,
		},
		Params: map[string]string{
			"invoiceId": invoiceID,
		},
		Response: &response,
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
