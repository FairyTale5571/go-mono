package acquiring

import (
	"context"
	"net/http"
	"time"

	"github.com/fairytale5571/go-mono/internal/api"
)

// MerchantWalletPaymentRequest запит оплати за токеном картки
type MerchantWalletPaymentRequest struct {
	CardToken        string            `json:"cardToken"`
	Amount           int               `json:"amount"`
	Ccy              int               `json:"ccy"`
	Tds              bool              `json:"tds"`
	RedirectURL      string            `json:"redirectUrl"`
	WebHookURL       string            `json:"webHookUrl"`
	InitiationKind   string            `json:"initiationKind"`
	MerchantPaymInfo *MerchantPaymInfo `json:"merchantPaymInfo"`
}

// MerchantWalletPaymentResponse відповідь на запит оплати за токеном картки
type MerchantWalletPaymentResponse struct {
	InvoiceID     string    `json:"invoiceId"`
	TdsURL        string    `json:"tdsUrl"`
	Status        string    `json:"status"`
	FailureReason string    `json:"failureReason"`
	Amount        int       `json:"amount"`
	Ccy           int       `json:"ccy"`
	CreatedDate   time.Time `json:"createdDate"`
	ModifiedDate  time.Time `json:"modifiedDate"`
}

// MerchantWalletPayment запит оплати за токеном картки
func (a *Acquiring) MerchantWalletPayment(ctx context.Context, req MerchantWalletPaymentRequest) (*MerchantWalletPaymentResponse, error) {
	var resp MerchantWalletPaymentResponse
	err := a.apiClient.SendRequest(ctx, api.Request{
		Method: http.MethodPost,
		Path:   "/wallet/payment",
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
