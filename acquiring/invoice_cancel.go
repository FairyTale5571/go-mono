package acquiring

import (
	"context"
	"net/http"
	"time"

	"github.com/fairytale5571/go-mono/internal/api"
)

// ItemCancel Список товарів для створення чеку повернення, поле обов'язкове у випадку активованої опції фіскалізації
type ItemCancel struct {
	Name    string  `json:"name"`
	Qty     float64 `json:"qty"`
	Sum     int64   `json:"sum"`
	Code    string  `json:"code"`
	Barcode string  `json:"barcode"`
	Header  string  `json:"header"`
	Footer  string  `json:"footer"`
	Tax     []int   `json:"tax"`
	UKTZED  string  `json:"uktzed"`
}

// InvoiceCancelRequest тіло запиту на скасування рахунку
type InvoiceCancelRequest struct {
	InvoiceID string `json:"invoiceId"`
	ExtRef    string `json:"extRef"`
	Amount    int64  `json:"amount"`
}

// NewDefaultInvoiceCancelRequest створення стандартного запиту на скасування рахунку
func (a *Acquiring) NewDefaultInvoiceCancelRequest() *InvoiceCancelRequest {
	return &InvoiceCancelRequest{}
}

// InvoiceCancelResponse тіло відповіді на скасування рахунку
type InvoiceCancelResponse struct {
	// Status Статус запиту - Enum: "processing" "success" "failure"
	Status       string    `json:"status"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

// CancelInvoice Скасування рахунку
func (a *Acquiring) CancelInvoice(ctx context.Context, req *InvoiceCancelRequest) (*InvoiceCancelResponse, error) {
	var resp InvoiceCancelResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/invoice/cancel",
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
