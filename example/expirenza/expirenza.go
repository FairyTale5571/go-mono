package main

import (
	"context"
	"github.com/FairyTale5571/go-mono/expirenza"
)

const (
	restoID   = "your-resto-id"
	secretKey = "your-secret-key"
)

func main() {
	client := expirenza.NewExpirenzaClient(expirenza.Settings{
		RestoID:   restoID,
		SecretKey: secretKey,
		Debug:     false,
	})

	ctx := context.Background()

	client.AddHandler(TablesInfoEvent)
	client.AddHandler(HallPlansInfoEvent)
	client.AddHandler(GetBillEvent)
	client.AddHandler(PayBillRequest)
	client.AddHandler(CreateOrderOnTableEvent)
	client.AddHandler(CheckProductsRestrictionsEvent)
	client.AddHandler(SplitOrderEvent)
	client.AddHandler(CategoriesInfoEvent)
	client.AddHandler(CheckWorkEvent)
	client.AddHandler(SetSettingsEvent)
	client.AddHandler(StopListEvent)
	client.AddHandler(VersionInfoEvent)
	client.AddHandler(MenuInfoEvent)
	client.AddHandler(UserInfoEvent)

	if err := client.Open(ctx); err != nil {
		panic(err)
	}

	select {}
}
