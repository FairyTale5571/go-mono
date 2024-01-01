package acquiring

import (
	"context"
	"net/http"
	"time"

	"github.com/FairyTale5571/go-mono/internal/api"
)

type CardData struct {
	Pan string `json:"pan"`
	Exp string `json:"exp"`
	CVV int    `json:"cvv"`
}

// InvoicePaymentDirectRequest тіло запиту на створення рахунку
type InvoicePaymentDirectRequest struct {
	Amount           int               `json:"amount"`
	Ccy              int               `json:"ccy"`
	CardData         *CardData         `json:"cardData"`
	MerchantPaymInfo *MerchantPaymInfo `json:"merchantPaymInfo"`
	WebHookURL       string            `json:"webHookUrl"`
	RedirectURL      string            `json:"redirectUrl"`
	PaymentType      PaymentType       `json:"paymentType"`
	SaveCardData     *SaveCardData     `json:"saveCardData"`
	InitiationKind   string            `json:"initiationKind"`
	TDS              bool              `json:"tds"`
}

// InvoicePaymentDirectResponse тіло відповіді на створення рахунку
type InvoicePaymentDirectResponse struct {
	InvoiceID     string    `json:"invoiceId"`
	TdsURL        string    `json:"tdsUrl"`
	Status        string    `json:"status"`
	FailureReason string    `json:"failureReason"`
	Amount        int       `json:"amount"`
	Ccy           int       `json:"ccy"`
	CreatedDate   time.Time `json:"createdDate"`
	ModifiedDate  time.Time `json:"modifiedDate"`
}

// CreateInvoicePaymentDirect Створення рахунку для оплати
func (a *Acquiring) CreateInvoicePaymentDirect(ctx context.Context, req *InvoicePaymentDirectRequest) (*InvoicePaymentDirectResponse, error) {
	var resp InvoicePaymentDirectResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/payment-direct",
		Body:   req,
		Headers: map[string]string{
			"X-Token":       a.token,
			"X-Cms":         a.cmsName,
			"X-Cms-Version": a.cmsVersion,
		},
		Response: &resp,
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
