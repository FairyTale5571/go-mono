package acquiring

import (
	"context"
	"net/http"
	"time"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// CancelList Список прийнятих заявок на скасування
type CancelList struct {
	Status       string    `json:"status"`
	Amount       int       `json:"amount"`
	Ccy          int       `json:"ccy"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
	ApprovalCode string    `json:"approvalCode"`
	RRN          string    `json:"rrn"`
	ExtRef       string    `json:"extRef"`
}

// WalletData Параметри картки
type WalletData struct {
	CardToken string `json:"cardToken"`
	WalletID  string `json:"walletId"`
	Status    string `json:"status"`
}

// InvoiceStatusResponse Відповідь на запит статусу рахунку
type InvoiceStatusResponse struct {
	InvoiceID     string        `json:"invoiceId"`
	Status        string        `json:"status"`
	FailureReason string        `json:"failureReason"`
	ErrCode       string        `json:"errCode"`
	Amount        int           `json:"amount"`
	Ccy           int           `json:"ccy"`
	FinalAmount   int           `json:"finalAmount"`
	CreatedDate   time.Time     `json:"createdDate"`
	ModifiedDate  time.Time     `json:"modifiedDate"`
	Reference     string        `json:"reference"`
	CancelList    []*CancelList `json:"cancelList"`
	WalletData    *WalletData   `json:"walletData"`
}

// GetInvoiceStatus Отримати статус рахунку
func (a *Acquiring) GetInvoiceStatus(ctx context.Context, invoiceID string) (*InvoiceStatusResponse, error) {
	var resp InvoiceStatusResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/invoice/status",
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
