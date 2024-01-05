package expirenza

const interfaceEventType Operation = "__ANY__"

type interfaceEventHandler func(*Expirenza, any)

func (eh interfaceEventHandler) Type() Operation {
	return interfaceEventType
}

func (eh interfaceEventHandler) New() any {
	return nil
}

// Handle is the handler for an any event.
func (eh interfaceEventHandler) Handle(s *Expirenza, i any) {
	eh(s, i)
}

// getBillHandler
type getBillHandler func(*Expirenza, *GetBillRequest)

func (h getBillHandler) Type() Operation {
	return OperationGetBill
}

func (h getBillHandler) New() any {
	return &GetBillRequest{}
}

func (h getBillHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*GetBillRequest); ok {
		h(e, t)
	}
}

// payBillHandler
type payBillHandler func(*Expirenza, *PayBillRequest)

func (h payBillHandler) Type() Operation {
	return OperationPayBill
}

func (h payBillHandler) New() any {
	return &PayBillRequest{}
}

func (h payBillHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*PayBillRequest); ok {
		h(e, t)
	}
}

type createOrderOnTableHandler func(*Expirenza, *CreateOrderOnTableRequest)

func (h createOrderOnTableHandler) Type() Operation {
	return OperationCreateOrderOnTable
}

func (h createOrderOnTableHandler) New() any {
	return &CreateOrderOnTableRequest{}
}

func (h createOrderOnTableHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CreateOrderOnTableRequest); ok {
		h(e, t)
	}
}

type checkProductsRestrictionsHandler func(*Expirenza, *CheckProductsRestrictionsRequest)

func (h checkProductsRestrictionsHandler) Type() Operation {
	return OperationCheckProductsRestrictions
}

func (h checkProductsRestrictionsHandler) New() any {
	return &CheckProductsRestrictionsRequest{}
}

func (h checkProductsRestrictionsHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CheckProductsRestrictionsRequest); ok {
		h(e, t)
	}
}

type splitOrderHandler func(*Expirenza, *SplitOrderRequest)

func (h splitOrderHandler) Type() Operation {
	return OperationSplitOrder
}

func (h splitOrderHandler) New() any {
	return &SplitOrderRequest{}
}

func (h splitOrderHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*SplitOrderRequest); ok {
		h(e, t)
	}
}

type tablesInfoHandler func(*Expirenza, *TablesInfoRequest)

func (h tablesInfoHandler) Type() Operation {
	return OperationTablesInfo
}

func (h tablesInfoHandler) New() any {
	return &TablesInfoRequest{}
}

func (h tablesInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*TablesInfoRequest); ok {
		h(e, t)
	}
}

type usersInfoHandler func(*Expirenza, *UsersInfoRequest)

func (h usersInfoHandler) Type() Operation {
	return OperationUsersInfo
}

func (h usersInfoHandler) New() any {
	return &UsersInfoRequest{}
}

func (h usersInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*UsersInfoRequest); ok {
		h(e, t)
	}
}

type categoriesInfoHandler func(*Expirenza, *CategoriesInfoRequest)

func (h categoriesInfoHandler) Type() Operation {
	return OperationCategoriesInfo
}

func (h categoriesInfoHandler) New() any {
	return &CategoriesInfoRequest{}
}

func (h categoriesInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CategoriesInfoRequest); ok {
		h(e, t)
	}
}

type hallPlansInfoHandler func(*Expirenza, *HallPlansInfoRequest)

func (h hallPlansInfoHandler) Type() Operation {
	return OperationHallPlansInfo
}

func (h hallPlansInfoHandler) New() any {
	return &HallPlansInfoRequest{}
}

func (h hallPlansInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*HallPlansInfoRequest); ok {
		h(e, t)
	}
}

type checkWorkHandler func(*Expirenza, *CheckWorkRequest)

func (h checkWorkHandler) Type() Operation {
	return OperationCheckWork
}

func (h checkWorkHandler) New() any {
	return &CheckWorkRequest{}
}

func (h checkWorkHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CheckWorkRequest); ok {
		h(e, t)
	}
}

type setSettingsHandler func(*Expirenza, *SetSettingsRequest)

func (h setSettingsHandler) Type() Operation {
	return OperationSetSettings
}

func (h setSettingsHandler) New() any {
	return &SetSettingsRequest{}
}

func (h setSettingsHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*SetSettingsRequest); ok {
		h(e, t)
	}
}

type stopListHandler func(*Expirenza, *StopListRequest)

func (h stopListHandler) Type() Operation {
	return OperationStopList
}

func (h stopListHandler) New() any {
	return &StopListRequest{}
}

func (h stopListHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*StopListRequest); ok {
		h(e, t)
	}
}

type versionInfoHandler func(*Expirenza, *VersionInfoRequest)

func (h versionInfoHandler) Type() Operation {
	return OperationVersionInfo
}

func (h versionInfoHandler) New() any {
	return &VersionInfoRequest{}
}

func (h versionInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*VersionInfoRequest); ok {
		h(e, t)
	}
}

type menuInfoHandler func(*Expirenza, *MenuInfoRequest)

func (h menuInfoHandler) Type() Operation {
	return OperationMenuInfo
}

func (h menuInfoHandler) New() any {
	return &MenuInfoRequest{}
}

func (h menuInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*MenuInfoRequest); ok {
		h(e, t)
	}
}
