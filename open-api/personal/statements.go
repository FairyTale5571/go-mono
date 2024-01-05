package personal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
)

// Statement - Об'єкт виписки
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
	ReceiptID       string `json:"receiptId"`
	InvoiceID       string `json:"invoiceId"`
	CounterEDRPOU   string `json:"counterEdrpou"`
	CounterIBAN     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}

// GetStatementsRequest - Запит на отримання виписки
type GetStatementsRequest struct {
	From    int64
	To      int64
	Account string
}

// GetStatements - Отримання виписки за час від from до to часу в секундах в форматі Unix time. Максимальний час, за який можливо отримати виписку — 31 доба + 1 година (2682000 секунд).
// Обмеження на використання функції — не частіше ніж 1 раз на 60 секунд.
// Повертає 500 транзакцій з кінця, тобто від часу to до from. Якщо кількість транзакцій = 500, потрібно зробити ще один запит, зменшивши час to до часу останнього платежу, з відповіді. Якщо знову кількість транзакцій = 500, то виконуєте запити до того часу, поки кількість транзакцій не буде < 500. Відповідно, якщо кількість транзакцій < 500, то вже отримано всі платежі за вказаний період.
func (p *Personal) GetStatements(ctx context.Context, req *GetStatementsRequest) ([]*Statement, error) {
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
