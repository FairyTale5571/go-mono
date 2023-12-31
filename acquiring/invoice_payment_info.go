package acquiring

import (
	"context"
	"github.com/fairytale5571/go-mono/internal/api"
	"net/http"
	"time"
)

type InvoiceInfoResponse struct {
	MaskedPan     string        `json:"maskedPan"`
	ApprovalCode  string        `json:"approvalCode"`
	RRN           string        `json:"rrn"`
	Amount        int           `json:"amount"`
	Ccy           int           `json:"ccy"`
	FinalAmount   int           `json:"finalAmount"`
	CreatedDate   time.Time     `json:"createdDate"`
	Terminal      string        `json:"terminal"`
	PaymentScheme string        `json:"paymentScheme"`
	PaymentMethod string        `json:"paymentMethod"`
	Fee           int           `json:"fee"`
	DomesticCard  bool          `json:"domesticCard"`
	Country       string        `json:"country"`
	CancelList    []*CancelList `json:"cancelList"`
}

// GetInvoiceInfo Дані про успішну оплату, якщо вона була здійснена
func (a *Acquiring) GetInvoiceInfo(ctx context.Context, invoiceID string) (*InvoiceInfoResponse, error) {
	var resp InvoiceInfoResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/invoice/payment-info",
		Params: map[string]string{
			"invoiceId": invoiceID,
		},
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
