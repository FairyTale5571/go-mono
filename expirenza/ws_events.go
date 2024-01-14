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

// GetBillEvent - Запит на отримання рахунку
type GetBillEvent struct {
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

// CreateOrderOnTableEvent - Запит на створення замовлення на столі
type CreateOrderOnTableEvent struct {
	*BaseEvent

	ExternalID string        `json:"externalId"`
	OriginName string        `json:"originName"`
	TableID    string        `json:"tableId"`
	Items      []*Item       `json:"items"`
	FiscalData []*FiscalData `json:"fiscalData"`
}

// CheckProductsRestrictionsEvent - Запит на перевірку наявності блюд
type CheckProductsRestrictionsEvent struct {
	*BaseEvent

	Items []*Item `json:"items"`
}

// SplitOrderEvent - Запит на розділення рахунку
type SplitOrderEvent struct {
	*BaseEvent

	BillID     string      `json:"billId"`
	ControlSum float64     `json:"controlSum"`
	Positions  []*Position `json:"positions"`
}

// TablesInfoEvent - Запит на отримання інформації про столи ресторану
type TablesInfoEvent struct {
	*BaseEvent
}

// UsersInfoEvent - Запит на отримання списку персоналу
type UsersInfoEvent struct {
	*BaseEvent
}

// CategoriesInfoEvent - Запит на отримання списку категорій
type CategoriesInfoEvent struct {
	*BaseEvent
}

// HallPlansInfoEvent - Запит на отримання переліку залів у ресторані
type HallPlansInfoEvent struct {
	*BaseEvent
}

// CheckWorkEvent - Запит на перевірку роботи ресторану
type CheckWorkEvent struct {
	*BaseEvent
}

// SetSettingsEvent - Запит на передачу налаштувань
type SetSettingsEvent struct {
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

// StopListEvent - Запит на отримання відповіді від ресторану про продукти, що знаходяться в стоп-листі
type StopListEvent struct {
	*BaseEvent
}

// VersionInfoEvent - Запит на отримання версії плагіну від ресторану
type VersionInfoEvent struct {
	*BaseEvent
}

// MenuInfoEvent - Запит на отримання меню ресторану
type MenuInfoEvent struct {
	*BaseEvent
}
