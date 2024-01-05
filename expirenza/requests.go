package expirenza

type Operation string

const (
	// OperationGetBill - Отримання рахунку за номером стола
	OperationGetBill Operation = "getBill"
	// OperationPayBill - Оплата рахунку і закриття замовлення
	OperationPayBill Operation = "payBill"
	// OperationCreateOrderOnTable - Створення замовлення на столі
	OperationCreateOrderOnTable Operation = "createOrderOnTable"
	// OperationCheckProductsRestrictions - Перевірка наявності блюд
	OperationCheckProductsRestrictions Operation = "checkProductsRestrictions"
	// OperationSplitOrder - Розділення рахунку (Split Bill)
	OperationSplitOrder Operation = "splitOrder"
	// OperationTablesInfo - Отримання інформації про столи ресторану
	OperationTablesInfo Operation = "tablesInfo"
	// OperationUsersInfo - Отримання списку персоналу
	OperationUsersInfo Operation = "usersInfo"
	// OperationCategoriesInfo - Отримання списку категорій
	OperationCategoriesInfo Operation = "categoriesInfo"
	// OperationHallPlansInfo - Отримання переліку залів у ресторані
	OperationHallPlansInfo Operation = "hallplansInfo"
	// OperationCheckWork - Перевірка роботи ресторану
	OperationCheckWork Operation = "checkWork"
	// OperationSetSettings - Передача налаштувань
	OperationSetSettings Operation = "setSettings"
	// OperationStopList - отримання відповіді від ресторану про продукти, що знаходяться в стоп-листі
	OperationStopList Operation = "stopList"
	// OperationVersionInfo отримання версії плагіну від ресторану
	OperationVersionInfo Operation = "versionInfo"
	// OperationMenuInfo - Отримання меню ресторану
	OperationMenuInfo Operation = "menuInfo"
)

// BaseEvent - Базовий запит
type BaseEvent struct {
	RID       string    `json:"rID"`
	Operation Operation `json:"operation"`
}

// GetOperation - Отримання типу запиту
func (e *BaseEvent) GetOperation() Operation {
	return e.Operation
}

// GetRID - Отримання ідентифікатора запиту
func (e *BaseEvent) GetRID() string {
	return e.RID
}

// GetBillRequest - Запит на отримання рахунку
type GetBillRequest struct {
	*BaseEvent

	BillID      string `json:"billId,omitempty"`
	TableNumber string `json:"tableNumber,omitempty"`
}

type FiscalData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PayBillRequest - Запит на оплату рахунку
type PayBillRequest struct {
	*BaseEvent

	BillID     string        `json:"billId,omitempty"`
	Sum        float64       `json:"sum"`
	Tips       float64       `json:"tips,omitempty"`
	FiscalData []*FiscalData `json:"fiscalData"`
}

// CreateOrderOnTableRequest - Запит на створення замовлення на столі
type CreateOrderOnTableRequest struct {
	*BaseEvent

	ExternalID string        `json:"externalId"`
	OriginName string        `json:"originName"`
	TableID    string        `json:"tableId"`
	Items      []*Item       `json:"items"`
	FiscalData []*FiscalData `json:"fiscalData"`
}

// CheckProductsRestrictionsRequest - Запит на перевірку наявності блюд
type CheckProductsRestrictionsRequest struct {
	*BaseEvent

	Items []*Item `json:"items"`
}

// SplitOrderRequest - Запит на розділення рахунку
type SplitOrderRequest struct {
	*BaseEvent

	BillID     string      `json:"billId"`
	ControlSum float64     `json:"controlSum"`
	Positions  []*Position `json:"positions"`
}

// TablesInfoRequest - Запит на отримання інформації про столи ресторану
type TablesInfoRequest struct {
	*BaseEvent
}

// UsersInfoRequest - Запит на отримання списку персоналу
type UsersInfoRequest struct {
	*BaseEvent
}

// CategoriesInfoRequest - Запит на отримання списку категорій
type CategoriesInfoRequest struct {
	*BaseEvent
}

// HallPlansInfoRequest - Запит на отримання переліку залів у ресторані
type HallPlansInfoRequest struct {
	*BaseEvent
}

// CheckWorkRequest - Запит на перевірку роботи ресторану
type CheckWorkRequest struct {
	*BaseEvent
}

// SetSettingsRequest - Запит на передачу налаштувань
type SetSettingsRequest struct {
	*BaseEvent

	PrintPreCheckQr         bool          `json:"printPrecheckQr"`
	PreCheckQrURL           string        `json:"precheckQrUrl"`
	MessageBeforeQr         string        `json:"messageBeforeQr"`
	MessageAfterQr          string        `json:"messageAfterQr"`
	PrintDeliveryPreCheckQr bool          `json:"printDeliveryPrecheckQr"`
	Settings                Modifiers     `json:"settings"`
	BonusPaymentTypes       []interface{} `json:"-"`
	BonusUsage              bool          `json:"bonusUsage"`
	PaymentAllowed          bool          `json:"paymentAllowed"`
}

// StopListRequest - Запит на отримання відповіді від ресторану про продукти, що знаходяться в стоп-листі
type StopListRequest struct {
	*BaseEvent
}

// VersionInfoRequest - Запит на отримання версії плагіну від ресторану
type VersionInfoRequest struct {
	*BaseEvent
}

// MenuInfoRequest - Запит на отримання меню ресторану
type MenuInfoRequest struct {
	*BaseEvent
}
