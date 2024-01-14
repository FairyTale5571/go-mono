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

func sign(secretKey string, body any) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	bytes, _ := json.Marshal(body)
	h.Write(bytes)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// HallPlansInfoResponse is a response to the HallPlansInfo method.
type HallPlansInfoResponse struct {
	HallPlans []*HallPlan `json:"hallPlans"`
	RID       string      `json:"rID"`
}

func (h *HallPlansInfoResponse) sign(secretKey string) string {
	return sign(secretKey, h)
}

// HallPlansInfo - відправляє інформацію про план залу.
func (e *Expirenza) HallPlansInfo(ctx context.Context, h *HallPlansInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/hallplansInfo",
			Body:   h,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": h.sign(e.secretKey),
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

func (t *TablesInfoResponse) sign(secretKey string) string {
	return sign(secretKey, t)
}

// TablesInfo - відправляє інформацію про столи.
func (e *Expirenza) TablesInfo(ctx context.Context, t *TablesInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/tablesInfo",
			Body:   t,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": t.sign(e.secretKey),
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

func (u *UsersInfoResponse) sign(secretKey string) string {
	return sign(secretKey, u)
}

// UsersInfo - відправляє інформацію про персонал.
func (e *Expirenza) UsersInfo(ctx context.Context, u *UsersInfoResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/usersInfo",
			Body:   u,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": u.sign(e.secretKey),
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

func (c *CheckWorkResponse) sign(secretKey string) string {
	return sign(secretKey, c)
}

// CheckWork - відправляє інформацію про готовність до роботи.
func (e *Expirenza) CheckWork(ctx context.Context, c CheckWorkResponse) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/checkWork",
			Body:   c,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": c.sign(e.secretKey),
			},
		},
	)
}

type GetBillResponse struct {
	ErrorCode    string  `json:"errorCode"`
	ErrorMessage string  `json:"errorMessage"`
	Orders       []Order `json:"orders"`
	RID          string  `json:"rID"`
	State        string  `json:"state"`
	SubState     string  `json:"substate"`
	TableNumber  int     `json:"tableNumber"`
}

func (g *GetBillResponse) sign(secretKey string) string {
	return sign(secretKey, g)
}

// GetBill - відправляє інформацію про рахунок.
func (e *Expirenza) GetBill(ctx context.Context, g *GetBillResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/getBill",
			Body:   g,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": g.sign(e.secretKey),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

type PayBillResponse struct {
	BillID            string `json:"billId"`
	Guests            Guest  `json:"guests"`
	OrderNumber       int    `json:"orderNumber"`
	RID               string `json:"rID"`
	StackTraceMessage string `json:"stackTraceMessage"`
	State             string `json:"state"`
	SubState          string `json:"substate"`
	Waiter            string `json:"waiter"`
	WaiterID          string `json:"waiterId"`
}

func (p *PayBillResponse) sign(secretKey string) string {
	return sign(secretKey, p)
}

// PayBill - відправляє інформацію про оплату рахунку.
func (e *Expirenza) PayBill(ctx context.Context, p *PayBillResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/payBill",
			Body:   p,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": p.sign(e.secretKey),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

type CreateOrderOnTableRequest struct {
	Operation  string       `json:"operation"`
	RID        string       `json:"rID"`
	ExternalID string       `json:"externalId"`
	OriginName string       `json:"originName"`
	TableID    string       `json:"tableId"`
	Items      []Item       `json:"items"`
	FiscalData []FiscalData `json:"fiscalData"`
}

func (c *CreateOrderOnTableRequest) sign(secretKey string) string {
	return sign(secretKey, c)
}

type CreateOrderOnTableResult struct {
	Success       bool           `json:"success"`
	Stage         string         `json:"stage"`
	RootOrder     RootOrder      `json:"rootOrder"`
	OrdersSummary []OrderSummary `json:"ordersSummary"`
	Guests        []Guest        `json:"guests"`
}

type CreateOrderOnTableResponse struct {
	RID    string                   `json:"rID"`
	Result CreateOrderOnTableResult `json:"result"`
}

// CreateOrderOnTable - Створення та оплата замовлення на столі
func (e *Expirenza) CreateOrderOnTable(ctx context.Context, c *CreateOrderOnTableRequest) (*CreateOrderOnTableResponse, error) {
	var result CreateOrderOnTableResponse
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/createOrderOnTable",
			Body:   c,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": c.sign(e.secretKey),
			},
			Response: &result,
		},
	); err != nil {
		return nil, err
	}
	return &result, nil
}

type CheckProductsRestrictionsRequest struct {
	Operation string `json:"operation"`
	RID       string `json:"rID"`
	Items     []Item `json:"items"`
}

func (c *CheckProductsRestrictionsRequest) sign(secretKey string) string {
	return sign(secretKey, c)
}

type CheckProductsRestrictionsResult struct {
	CheckResult string       `json:"checkResult"`
	ItemResults []ItemResult `json:"itemResults"`
}

type CheckProductsRestrictionsResponse struct {
	RID    string                          `json:"rID"`
	Result CheckProductsRestrictionsResult `json:"result"`
}

// CheckProductsRestrictions - Перевірка обмежень на продукти
func (e *Expirenza) CheckProductsRestrictions(ctx context.Context, c *CheckProductsRestrictionsRequest) (*CheckProductsRestrictionsResponse, error) {
	var result CheckProductsRestrictionsResponse
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/checkProductsRestrictions",
			Body:   c,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": c.sign(e.secretKey),
			},
			Response: &result,
		},
	); err != nil {
		return nil, err
	}
	return &result, nil
}

type SplitOrderResponse struct {
	RID         string  `json:"rID"`
	State       string  `json:"state"`
	TableNumber int     `json:"tableNumber"`
	Orders      []Order `json:"orders"`
}

func (s *SplitOrderResponse) sign(secretKey string) string {
	return sign(secretKey, s)
}

func (e *Expirenza) SplitOrder(ctx context.Context, s *SplitOrderResponse) error {
	if err := e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/splitOrder",
			Body:   s,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": s.sign(e.secretKey),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

type CategoriesInfoResponse struct {
	RID        string      `json:"rID"`
	Categories []*Category `json:"categories"`
}

func (c *CategoriesInfoResponse) sign(secretKey string) string {
	return sign(secretKey, c)
}

func (e *Expirenza) CategoriesInfo(ctx context.Context, c *CategoriesInfoResponse) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/categoriesInfo",
			Body:   c,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": c.sign(e.secretKey),
			},
		},
	)
}

type StopListResponse struct {
	RID   string `json:"rID"`
	Items []Item `json:"items"`
}

func (s *StopListResponse) sign(secretKey string) string {
	return sign(secretKey, s)
}

func (e *Expirenza) StopList(ctx context.Context, s *StopListResponse) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/stopList",
			Body:   s,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": s.sign(e.secretKey),
			},
		},
	)
}

type VersionInfoResponse struct {
	// AdditionalInfo - TODO: ?
	AdditionalInfo             any          `json:"additionalInfo"`
	PartnerID                  string       `json:"partner.id"`
	PluginCompilationTimestamp string       `json:"plugin.compilation.timestamp"`
	PluginVersion              string       `json:"plugin.version"`
	RID                        string       `json:"rID"`
	RestaurantLicenseDateEnd   string       `json:"restaurant.license.date.end"`
	RestaurantVersion          string       `json:"restaurant.version"`
	SystemInfo                 SystemInfo   `json:"systemInfo"`
	UpdateInfo                 []UpdateInfo `json:"updateInfo"`
}

func (v *VersionInfoResponse) sign(secretKey string) string {
	return sign(secretKey, v)
}

func (e *Expirenza) VersionInfo(ctx context.Context, v *VersionInfoResponse) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/versionInfo",
			Body:   v,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": v.sign(e.secretKey),
			},
		},
	)
}

type MenuInfoRequest struct {
	// Dishes - Страви
	Dishes []MenuDish `json:"dishes"`
	// GroupModifiers - групи модифікаторів страв
	GroupModifiers []GroupModifier `json:"groupModifiers"`
	// Groups - розділи меню
	Groups []Group `json:"groups"`
	// HostTerminalsGroup -
	HostTerminalsGroup TerminalGroup `json:"hostTerminalsGroup"`
	// Modifiers - модифікатори страв
	Modifiers []MenuModifier `json:"modifiers"`
	RID       string         `json:"rID"`
	// TerminalsGroups -
	TerminalsGroups []TerminalGroup `json:"terminalsGroups"`
}

func (m *MenuInfoRequest) sign(secretKey string) string {
	return sign(secretKey, m)
}

// MenuInfo - відправляє інформацію про меню.
func (e *Expirenza) MenuInfo(ctx context.Context, m *MenuInfoRequest) error {
	return e.apiClient.SendRequest(ctx,
		api.Request{
			Method: http.MethodPost,
			Path:   "/callback/menuInfo",
			Body:   m,
			Headers: map[string]string{
				"restoId":   e.restoID,
				"signature": m.sign(e.secretKey),
			},
		},
	)
}
