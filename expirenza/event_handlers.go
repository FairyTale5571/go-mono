package expirenza

func (e *Expirenza) handlerForInterface(handler interface{}) EventHandler {
	switch v := handler.(type) {
	case func(*Expirenza, interface{}):
		return interfaceEventHandler(v)
	case func(*Expirenza, *GetBillRequest):
		return getBillHandler(v)
	case func(*Expirenza, *PayBillRequest):
		return payBillHandler(v)
	case func(*Expirenza, *CreateOrderOnTableRequest):
		return createOrderOnTableHandler(v)
	case func(*Expirenza, *CheckProductsRestrictionsRequest):
		return checkProductsRestrictionsHandler(v)
	case func(*Expirenza, *SplitOrderRequest):
		return splitOrderHandler(v)
	case func(*Expirenza, *TablesInfoRequest):
		return tablesInfoHandler(v)
	case func(*Expirenza, *UsersInfoRequest):
		return usersInfoHandler(v)
	case func(*Expirenza, *CategoriesInfoRequest):
		return categoriesInfoHandler(v)
	case func(*Expirenza, *HallPlansInfoRequest):
		return hallPlansInfoHandler(v)
	case func(*Expirenza, *CheckWorkRequest):
		return checkWorkHandler(v)
	case func(*Expirenza, *SetSettingsRequest):
		return setSettingsHandler(v)
	case func(*Expirenza, *StopListRequest):
		return stopListHandler(v)
	case func(*Expirenza, *VersionInfoRequest):
		return versionInfoHandler(v)
	case func(*Expirenza, *MenuInfoRequest):
		return menuInfoHandler(v)
	default:
		return nil
	}
}

// nolint: gochecknoinits
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
