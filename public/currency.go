package public

import (
	"context"
	"net/http"

	"github.com/fairytale5571/go-mono/internal/api"
)

type Currency struct {
	// CurrencyCodeA - Код валюти рахунку відповідно ISO 4217
	CurrencyCodeA int32 `json:"currencyCodeA"`
	// CurrencyCodeB - Код валюти рахунку відповідно ISO 4217
	CurrencyCodeB int32 `json:"currencyCodeB"`
	// Date - Час курсу в секундах в форматі Unix time
	Date      int     `json:"date"`
	RateSell  float64 `json:"rateSell"`
	RateBuy   float64 `json:"rateBuy"`
	RateCross float64 `json:"rateCross"`
}

// GetCurrency - Отримати базовий перелік курсів валют monobank. Інформація кешується та оновлюється не частіше 1 разу на 5 хвилин.
func (p *Public) GetCurrencies(ctx context.Context) ([]*Currency, error) {
	var resp []*Currency
	err := p.apiClient.SendRequest(ctx, api.Request{
		Method:   http.MethodGet,
		Path:     "/bank/currency",
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
