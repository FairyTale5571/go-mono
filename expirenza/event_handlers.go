package expirenza

// you need fix this code gocyclo
func (e *Expirenza) handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(*Expirenza, interface{}):
		return interfaceEventHandler(v)
	case func(*Expirenza, *GetBillEvent):
		return getBillHandler(v)
	case func(*Expirenza, *PayBillRequest):
		return payBillHandler(v)
	case func(*Expirenza, *CreateOrderOnTableEvent):
		return createOrderOnTableHandler(v)
	case func(*Expirenza, *CheckProductsRestrictionsEvent):
		return checkProductsRestrictionsHandler(v)
	case func(*Expirenza, *SplitOrderEvent):
		return splitOrderHandler(v)
	case func(*Expirenza, *TablesInfoEvent):
		return tablesInfoHandler(v)
	case func(*Expirenza, *UsersInfoEvent):
		return usersInfoHandler(v)
	case func(*Expirenza, *CategoriesInfoEvent):
		return categoriesInfoHandler(v)
	case func(*Expirenza, *HallPlansInfoEvent):
		return hallPlansInfoHandler(v)
	case func(*Expirenza, *CheckWorkEvent):
		return checkWorkHandler(v)
	case func(*Expirenza, *SetSettingsEvent):
		return setSettingsHandler(v)
	case func(*Expirenza, *StopListEvent):
		return stopListHandler(v)
	case func(*Expirenza, *VersionInfoEvent):
		return versionInfoHandler(v)
	case func(*Expirenza, *MenuInfoEvent):
		return menuInfoHandler(v)
	default:
		return nil
	}
}

// nolint: gochecknoinits // this is init
func init() {
	registerInterfaceProvider(getBillHandler(nil))
	registerInterfaceProvider(payBillHandler(nil))
	registerInterfaceProvider(createOrderOnTableHandler(nil))
	registerInterfaceProvider(checkProductsRestrictionsHandler(nil))
	registerInterfaceProvider(splitOrderHandler(nil))
	registerInterfaceProvider(tablesInfoHandler(nil))
	registerInterfaceProvider(usersInfoHandler(nil))
	registerInterfaceProvider(categoriesInfoHandler(nil))
	registerInterfaceProvider(hallPlansInfoHandler(nil))
	registerInterfaceProvider(checkWorkHandler(nil))
	registerInterfaceProvider(setSettingsHandler(nil))
	registerInterfaceProvider(stopListHandler(nil))
	registerInterfaceProvider(versionInfoHandler(nil))
	registerInterfaceProvider(menuInfoHandler(nil))
}
