package chast

// PersonalDataRequest тіло запиту на отримання персональних даних
type PersonalDataRequest struct {
	CallbackURL string `json:"callbackUrl"`
	Extref      string `json:"extref"`
	Phone       string `json:"phone"`
}

// PersonalDataResponse тіло відповіді на отримання персональних даних
type PersonalDataResponse struct {
	Extref        string  `json:"extref"`
	ClientPhone   string  `json:"clientPhone,omitempty"`
	OrderSubState string  `json:"order_sub_state"`
	Client        *Client `json:"client,omitempty"`
}
