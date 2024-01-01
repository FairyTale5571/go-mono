package acquiring

import (
	"context"
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
	"github.com/FairyTale5571/go-mono/iso"
)

// InvoiceCreateRequest тіло запиту на створення рахунку
type InvoiceCreateRequest struct {
	Amount           int64             `json:"amount"`
	CCY              iso.ISO4217       `json:"ccy"`
	MerchantPaymInfo *MerchantPaymInfo `json:"merchantPaymInfo"`
	RedirectURL      string            `json:"redirectUrl"`
	WebHookURL       string            `json:"webHookUrl"`
	Validity         int               `json:"validity"`
	PaymentType      PaymentType       `json:"paymentType"`
	QrID             string            `json:"qrId"`
	Code             string            `json:"code"`
	SaveCardData     *SaveCardData     `json:"saveCardData"`
}

// NewDefaultInvoiceCreateRequest створення стандартного запиту на створення рахунку
func (a *Acquiring) NewDefaultInvoiceCreateRequest() *InvoiceCreateRequest {
	return &InvoiceCreateRequest{
		Validity:    86400,
		PaymentType: Debit,
		CCY:         iso.UAH,
	}
}

// InvoiceCreateResponse тіло відповіді на створення рахунку
type InvoiceCreateResponse struct {
	InvoiceID string `json:"invoiceId"`
	PageURL   string `json:"pageUrl"`
}

// CreateInvoice Створення рахунку для оплати
func (a *Acquiring) CreateInvoice(ctx context.Context, req *InvoiceCreateRequest) (*InvoiceCreateResponse, error) {
	var resp InvoiceCreateResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/create",
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
