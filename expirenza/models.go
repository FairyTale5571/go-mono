package expirenza

// HallPlan - обʼєкт залу
type HallPlan struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Table - обʼєкт столика
type Table struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	HallPlanID string `json:"hallPlanId"`
	Number     int    `json:"number"`
}

// User - обʼєкт користувача
type User struct {
	ID        string `json:"id"`
	CellPhone string `json:"cellPhone"`
	Name      string `json:"name"`
	Roles     string `json:"roles"`
}

// Modifier - обʼєкт модифікатора
type Modifier struct {
	ID       string `json:"id"`
	GroupID  string `json:"groupId"`
	Quantity int    `json:"quantity"`
}

// Item - обʼєкт в меню
type Item struct {
	ID       string      `json:"id"`
	Quantity float64     `json:"quantity"`
	Modifier []*Modifier `json:"modifier,omitempty"`
}

type Modifiers struct {
	LoadModifiers string `json:"load_modifiers"`
}

type Position struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

// Dish - обʼєкт страви
type Dish struct {
	Category      string  `json:"category"`
	CategoryID    string  `json:"categoryId"`
	CheckNumber   string  `json:"checkNumber"`
	Count         float64 `json:"count"`
	DishID        string  `json:"dishId"`
	LinkedTo      string  `json:"linkedTo"`
	Name          string  `json:"name"`
	OrderNumber   int     `json:"orderNumber"`
	PositionID    string  `json:"positionId"`
	Sum           float64 `json:"sum"`
	TaxCategory   string  `json:"taxCategory"`
	TaxCategoryID string  `json:"taxCategoryId"`
	Type          string  `json:"type"`
}

type Guest struct {
	Dishes []Dish `json:"dishes"`
	Name   string `json:"name"`
}

type List struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Sum  float64 `json:"sum"`
}

type Discount struct {
	List []List `json:"list"`
}

// Order - обʼєкт замовлення
type Order struct {
	BillID   string  `json:"billId"`
	BonusSum float64 `json:"bonusSum"`
	Created  string  `json:"created"`
	// Discounts - TODO: Ask expirenza developer what exactly fields in this struct object
	// Right now we don't know what fields in this struct object
	Discounts      Discount `json:"discounts"`
	Guests         []Guest  `json:"guests"`
	IsBanquetOrder bool     `json:"isBanquetOrder"`
	IsDelivery     bool     `json:"isDelivery"`
	OrderNumber    int      `json:"orderNumber"`
	PaymentType    string   `json:"paymentType"`
	PricingTime    string   `json:"pricingTime"`
	RawSum         int      `json:"rawSum"`
	RootBillID     string   `json:"rootBillId"`
	SplitAvailable bool     `json:"splitAvailable"`
	State          string   `json:"state"`
	TableID        string   `json:"tableId"`
	TotalSum       float64  `json:"totalSum"`
	Waiter         string   `json:"waiter"`
	WaiterID       string   `json:"waiterId"`
}

// OrderSummary - обʼєкт замовлення
type OrderSummary struct {
	OrderID     string  `json:"orderId"`
	OrderNumber int     `json:"orderNumber"`
	ResultSum   float64 `json:"resultSum"`
	OrderStatus string  `json:"orderStatus"`
	Closed      string  `json:"closed"`
	TableNumber int     `json:"tableNumber"`
	WaiterID    string  `json:"waiterId"`
	Waiter      string  `json:"waiter"`
}

type RootOrder struct {
	OrderID     string  `json:"orderId"`
	OrderNumber int     `json:"orderNumber"`
	ResultSum   float64 `json:"resultSum"`
	OrderStatus string  `json:"orderStatus"`
	TableNumber int     `json:"tableNumber"`
	WaiterID    string  `json:"waiterId"`
	Waiter      string  `json:"waiter"`
}

type ItemResult struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Available      bool    `json:"available"`
	ExceedQuantity float64 `json:"exceedQuantity"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type StopListItem struct {
	Balance   int    `json:"balance"`
	ProductID string `json:"productId"`
}

type SystemInfo struct {
	AskCashierForMultiCashRegisterPayment bool   `json:"AskCashierForMultiCashRegisterPayment"`
	HostTerminalName                      string `json:"HostTerminalName"`
	IsRoundingDownOrderSum                bool   `json:"IsRoundingDownOrderSum"`
	// AdditionalInfo - TODO: ?
	AdditionalInfo any  `json:"additionalInfo"`
	IsMain         bool `json:"isMain"`
}

type UpdateInfo struct {
	Enabled        bool   `json:"enabled"`
	LastRunTime    string `json:"lastRunTime"`
	LastTaskResult int    `json:"lastTaskResult"`
	TaskName       string `json:"taskName"`
}

// Types for MenuInfoEvent response
type (
	MenuDish struct {
		Category         string          `json:"category"`
		CategoryID       string          `json:"categoryId"`
		Description      string          `json:"description"`
		GroupID          string          `json:"groupId"`
		GroupModifiers   []GroupModifier `json:"groupModifiers"`
		ID               string          `json:"id"`
		IsActive         bool            `json:"isActive"`
		MenuSections     []string        `json:"menuSections"`
		ModifierSchemeID string          `json:"modifierSchemeId"`
		Modifiers        []Modifier      `json:"modifiers"`
		Name             string          `json:"name"`
		Price            float64         `json:"price"`
		Scale            Scale           `json:"scale"`
		TaxCategory      string          `json:"taxCategory"`
		TaxCategoryID    string          `json:"taxCategoryId"`
		Type             string          `json:"type"`
		Unit             string          `json:"unit"`
		VendorCode       string          `json:"vendorCode"`
	}

	GroupModifier struct {
		ChildModifiers                       []ChildModifier `json:"childModifiers"`
		ChildModifiersHaveMinMaxRestrictions bool            `json:"childModifiersHaveMinMaxRestrictions"`
		DefaultAmount                        int             `json:"defaultAmount"`
		Free                                 bool            `json:"free"`
		FreeAmount                           int             `json:"freeAmount"`
		ID                                   string          `json:"id"`
		MaxAmount                            int             `json:"maxAmount"`
		MinAmount                            int             `json:"minAmount"`
		Required                             bool            `json:"required"`
	}

	ChildModifier struct {
		DefaultAmount int    `json:"defaultAmount"`
		ID            string `json:"id"`
		MaxAmount     int    `json:"maxAmount"`
		MinAmount     int    `json:"minAmount"`
		Required      bool   `json:"required"`
	}

	Scale struct {
		DefaultSize DefaultSize `json:"defaultSize"`
		ID          string      `json:"id"`
		Name        string      `json:"name"`
	}

	DefaultSize struct {
		ID          string `json:"id"`
		KitchenName string `json:"kitchenName"`
		MenuIndex   int    `json:"menuIndex"`
		Name        string `json:"name"`
	}

	MenuModifier struct {
		Category      string  `json:"category"`
		CategoryID    string  `json:"categoryId"`
		Description   string  `json:"description"`
		GroupID       string  `json:"groupId"`
		ID            string  `json:"id"`
		IsActive      bool    `json:"isActive"`
		Name          string  `json:"name"`
		Price         float64 `json:"price"`
		Scale         Scale   `json:"scale"`
		TaxCategory   string  `json:"taxCategory"`
		TaxCategoryID string  `json:"taxCategoryId"`
		Unit          string  `json:"unit"`
		VendorCode    string  `json:"vendorCode"`
	}

	Group struct {
		ID        string `json:"id"`
		MenuIndex int    `json:"menuIndex"`
		Name      string `json:"name"`
		ParentID  string `json:"parentId"`
	}

	TerminalGroup struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
