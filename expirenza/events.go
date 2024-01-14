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
type getBillHandler func(*Expirenza, *GetBillEvent)

func (h getBillHandler) Type() Operation {
	return OperationGetBill
}

func (h getBillHandler) New() any {
	return &GetBillEvent{}
}

func (h getBillHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*GetBillEvent); ok {
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

type createOrderOnTableHandler func(*Expirenza, *CreateOrderOnTableEvent)

func (h createOrderOnTableHandler) Type() Operation {
	return OperationCreateOrderOnTable
}

func (h createOrderOnTableHandler) New() any {
	return &CreateOrderOnTableEvent{}
}

func (h createOrderOnTableHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CreateOrderOnTableEvent); ok {
		h(e, t)
	}
}

type checkProductsRestrictionsHandler func(*Expirenza, *CheckProductsRestrictionsEvent)

func (h checkProductsRestrictionsHandler) Type() Operation {
	return OperationCheckProductsRestrictions
}

func (h checkProductsRestrictionsHandler) New() any {
	return &CheckProductsRestrictionsEvent{}
}

func (h checkProductsRestrictionsHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CheckProductsRestrictionsEvent); ok {
		h(e, t)
	}
}

type splitOrderHandler func(*Expirenza, *SplitOrderEvent)

func (h splitOrderHandler) Type() Operation {
	return OperationSplitOrder
}

func (h splitOrderHandler) New() any {
	return &SplitOrderEvent{}
}

func (h splitOrderHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*SplitOrderEvent); ok {
		h(e, t)
	}
}

type tablesInfoHandler func(*Expirenza, *TablesInfoEvent)

func (h tablesInfoHandler) Type() Operation {
	return OperationTablesInfo
}

func (h tablesInfoHandler) New() any {
	return &TablesInfoEvent{}
}

func (h tablesInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*TablesInfoEvent); ok {
		h(e, t)
	}
}

type usersInfoHandler func(*Expirenza, *UsersInfoEvent)

func (h usersInfoHandler) Type() Operation {
	return OperationUsersInfo
}

func (h usersInfoHandler) New() any {
	return &UsersInfoEvent{}
}

func (h usersInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*UsersInfoEvent); ok {
		h(e, t)
	}
}

type categoriesInfoHandler func(*Expirenza, *CategoriesInfoEvent)

func (h categoriesInfoHandler) Type() Operation {
	return OperationCategoriesInfo
}

func (h categoriesInfoHandler) New() any {
	return &CategoriesInfoEvent{}
}

func (h categoriesInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CategoriesInfoEvent); ok {
		h(e, t)
	}
}

type hallPlansInfoHandler func(*Expirenza, *HallPlansInfoEvent)

func (h hallPlansInfoHandler) Type() Operation {
	return OperationHallPlansInfo
}

func (h hallPlansInfoHandler) New() any {
	return &HallPlansInfoEvent{}
}

func (h hallPlansInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*HallPlansInfoEvent); ok {
		h(e, t)
	}
}

type checkWorkHandler func(*Expirenza, *CheckWorkEvent)

func (h checkWorkHandler) Type() Operation {
	return OperationCheckWork
}

func (h checkWorkHandler) New() any {
	return &CheckWorkEvent{}
}

func (h checkWorkHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*CheckWorkEvent); ok {
		h(e, t)
	}
}

type setSettingsHandler func(*Expirenza, *SetSettingsEvent)

func (h setSettingsHandler) Type() Operation {
	return OperationSetSettings
}

func (h setSettingsHandler) New() any {
	return &SetSettingsEvent{}
}

func (h setSettingsHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*SetSettingsEvent); ok {
		h(e, t)
	}
}

type stopListHandler func(*Expirenza, *StopListEvent)

func (h stopListHandler) Type() Operation {
	return OperationStopList
}

func (h stopListHandler) New() any {
	return &StopListEvent{}
}

func (h stopListHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*StopListEvent); ok {
		h(e, t)
	}
}

type versionInfoHandler func(*Expirenza, *VersionInfoEvent)

func (h versionInfoHandler) Type() Operation {
	return OperationVersionInfo
}

func (h versionInfoHandler) New() any {
	return &VersionInfoEvent{}
}

func (h versionInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*VersionInfoEvent); ok {
		h(e, t)
	}
}

type menuInfoHandler func(*Expirenza, *MenuInfoEvent)

func (h menuInfoHandler) Type() Operation {
	return OperationMenuInfo
}

func (h menuInfoHandler) New() any {
	return &MenuInfoEvent{}
}

func (h menuInfoHandler) Handle(e *Expirenza, i any) {
	if t, ok := i.(*MenuInfoEvent); ok {
		h(e, t)
	}
}
