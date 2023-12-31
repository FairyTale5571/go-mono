package acquiring

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/fairytale5571/go-mono/internal/api"
)

// InvoiceStatementRequest відповідь на запит статусу рахунку
type InvoiceStatementRequest struct {
	// From - початкова дата в форматі UTC Unix Timestamp
	From int64
	// To - кінцева дата в форматі UTC Unix Timestamp
	To int64
}

// InvoiceStatement відповідь на запит статусу рахунку
type InvoiceStatement struct {
	InvoiceID     string        `json:"invoiceId"`
	Status        string        `json:"status"`
	MaskedPan     string        `json:"maskedPan"`
	Date          time.Time     `json:"date"`
	PaymentScheme string        `json:"paymentScheme"`
	Amount        int           `json:"amount"`
	ProfitAmount  int           `json:"profitAmount"`
	Ccy           int           `json:"ccy"`
	ApprovalCode  string        `json:"approvalCode"`
	Rrn           string        `json:"rrn"`
	Reference     string        `json:"reference"`
	ShortQrID     string        `json:"shortQrId"`
	CancelList    []*CancelList `json:"cancelList"`
}

// InvoiceStatementResponse відповідь на запит статусу рахунку
type InvoiceStatementResponse struct {
	List []*InvoiceStatement `json:"list"`
}

// InvoiceStatement відповідь на запит статусу рахунку
func (a *Acquiring) InvoiceStatement(ctx context.Context, req InvoiceStatementRequest) (*InvoiceStatementResponse, error) {
	var resp InvoiceStatementResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodGet,
		Path:   "/statement",
		Params: map[string]string{
			"from": strconv.FormatInt(req.From, 10),
			"to":   strconv.FormatInt(req.To, 10),
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
