package public

import (
	"context"
	"fmt"
	"github.com/fairytale5571/go-mono/internal/api"
	"net/http"
)

type Statement struct {
	ID              string `json:"id"`
	Time            int    `json:"time"`
	Description     string `json:"description"`
	Mcc             int    `json:"mcc"`
	OriginalMcc     int    `json:"originalMcc"`
	Hold            bool   `json:"hold"`
	Amount          int    `json:"amount"`
	OperationAmount int    `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int    `json:"commissionRate"`
	CashbackAmount  int    `json:"cashbackAmount"`
	Balance         int    `json:"balance"`
	Comment         string `json:"comment"`
	ReceiptId       string `json:"receiptId"`
	InvoiceId       string `json:"invoiceId"`
	CounterEDRPOU   string `json:"counterEdrpou"`
	CounterIBAN     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}

type GetStatementsRequest struct {
	From    int64
	To      int64
	Account string
}

// GetStatements - Отримання виписки за час від from до to часу в секундах в форматі Unix time. Максимальний час, за який можливо отримати виписку — 31 доба + 1 година (2682000 секунд).
// Обмеження на використання функції — не частіше ніж 1 раз на 60 секунд.
func (p *Public) GetStatements(ctx context.Context, req *GetStatementsRequest) ([]*Statement, error) {
	var resp []*Statement

	err := p.apiClient.SendRequest(ctx, api.Request{
		Method:   http.MethodGet,
		Path:     fmt.Sprintf("/personal/statement/%s/%d/%d", req.Account, req.From, req.To),
		Response: &resp,
		Headers: map[string]string{
			"X-Token": p.token,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
