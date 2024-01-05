package expirenza

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/FairyTale5571/go-mono/internal/api"
)

type HallPlansInfoResponse struct {
	HallPlans []*HallPlan `json:"hallPlans"`
	RID       string      `json:"rID"`
}

func sign(secretKey string, body interface{}) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	bytes, _ := json.Marshal(body)
	h.Write(bytes)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (h *HallPlansInfoResponse) Sign(secretKey string) string {
	return ""
}

func (e *Expirenza) HallPlansInfo(ctx context.Context, h *HallPlansInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/hallPlansInfo",
			Body:   h,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": h.Sign(e.secretKey),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

type TablesInfoResponse struct {
	RID    string   `json:"rID"`
	Tables []*Table `json:"tables"`
}

func (t *TablesInfoResponse) Sign(secretKey string) string {
	return sign(secretKey, t)
}

func (e *Expirenza) TablesInfo(ctx context.Context, t *TablesInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/tablesInfo",
			Body:   t,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": t.Sign(e.secretKey),
			},
			Response: nil,
		},
	); err != nil {
		return err
	}
	return nil
}

type UsersInfoResponse struct {
	RID   string  `json:"rID"`
	Users []*User `json:"users"`
}

func (u *UsersInfoResponse) Sign(secretKey string) string {
	return sign(secretKey, u)
}

func (e *Expirenza) UsersInfo(ctx context.Context, u *UsersInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/usersInfo",
			Body:   u,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": u.Sign(e.secretKey),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

type CheckWorkResponse struct {
	DateTime    string `json:"dateTime"`
	RID         string `json:"rID"`
	ReadyToWork bool   `json:"readyToWork"`
}

func (c *CheckWorkResponse) Sign(secretKey string) string {
	return sign(secretKey, c)
}

func (e *Expirenza) CheckWork(ctx context.Context, c CheckWorkResponse) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/checkWork",
			Body:   c,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": "",
			},
		},
	)
}
