package public

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
	"github.com/shopspring/decimal"
)

// Account - Рахунки клієнта
type Account struct {
	ID           string          `json:"id"`
	SendID       string          `json:"sendId"`
	Balance      decimal.Decimal `json:"balance"`
	CreditLimit  decimal.Decimal `json:"creditLimit"`
	Type         string          `json:"type"`
	CurrencyCode int             `json:"currencyCode"`
	CashbackType string          `json:"cashbackType"`
	MaskedPan    []string        `json:"maskedPan"`
	Iban         string          `json:"iban"`
}

// Jar - Банки клієнта
type Jar struct {
	Id           string          `json:"id"`
	SendID       string          `json:"sendId"`
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	CurrencyCode int             `json:"currencyCode"`
	Balance      decimal.Decimal `json:"balance"`
	Goal         decimal.Decimal `json:"goal"`
}

// InfoClient - Інформація про клієнта та переліку його рахунків і банок
type InfoClient struct {
	// ClientID - Ідентифікатор клієнта (збігається з id для send.monobank.ua)
	ClientID string `json:"clientId"`
	// Name - Ім'я клієнта
	Name string `json:"name"`
	// WebHookURL - URL для надсиляння подій по зміні балансу рахунку
	WebHookURL string `json:"webHookUrl"`
	// Permissions - Перелік прав, які які надає сервіс (1 літера на 1 permission).
	Permissions string `json:"permissions"`
	// Accounts - Перелік доступних рахунків
	Accounts []*Account `json:"accounts"`
	// Jars - Перелік доступних банок
	Jars []*Jar `json:"jars"`
}

// GetInfoClient - Отримання інформації про клієнта та переліку його рахунків і банок. Обмеження на використання функції не частіше ніж 1 раз у 60 секунд.
func (p *Public) GetInfoClient(ctx context.Context) (*InfoClient, error) {
	var resp InfoClient
	err := p.apiClient.SendRequest(ctx, api.Request{
		Method:   http.MethodGet,
		Path:     "/personal/client-info",
		Response: &resp,
		Headers: map[string]string{
			"X-Token": p.token,
		},
	})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
