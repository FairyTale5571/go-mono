package acquiring

// MerchantPaymInfo - інформація про платіж
type MerchantPaymInfo struct {
	Reference      string         `json:"reference"`
	Destination    string         `json:"destination"`
	Comment        string         `json:"comment"`
	CustomerEmails []string       `json:"customerEmails"`
	BasketOrder    []*BasketOrder `json:"basketOrder"`
}

// BasketOrder - товари в чеку
type BasketOrder struct {
	Name      string      `json:"name"`
	Qty       float64     `json:"qty"`
	Sum       int64       `json:"sum"`
	Icon      string      `json:"icon"`
	Unit      string      `json:"unit"`
	Code      string      `json:"code"`
	Barcode   string      `json:"barcode"`
	Header    string      `json:"header"`
	Footer    string      `json:"footer"`
	Tax       []int       `json:"tax"`
	UKTZED    string      `json:"uktzed"`
	Discounts []*Discount `json:"discounts"`
}

type DiscountType string

const (
	// DiscountTypeDiscount - знижка
	DiscountTypeDiscount DiscountType = "DISCOUNT"

	// DiscountTypeExtraCharge - надбавка
	DiscountTypeExtraCharge DiscountType = "EXTRA_CHARGE"
)

type DiscountMode string

const (
	// DiscountModePercent - відсоток
	DiscountModePercent DiscountMode = "PERCENT"

	// DiscountModeValue - сума
	DiscountModeValue DiscountMode = "VALUE"
)

type Discount struct {
	Type  DiscountType `json:"type"`
	Mode  DiscountMode `json:"mode"`
	Value float64      `json:"value"`
}

type SaveCardData struct {
	SaveCard bool   `json:"saveCard"`
	WalletID string `json:"walletId"`
}

type PaymentType string

const (
	Debit PaymentType = "debit"
	Hold  PaymentType = "hold"
)
