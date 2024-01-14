package main

import (
	"context"
	"fmt"

	"github.com/FairyTale5571/go-mono/expirenza"
)

func UserInfoEvent(e *expirenza.Expirenza, s *expirenza.UsersInfoEvent) {
	if err := e.UsersInfo(context.Background(), &expirenza.UsersInfoResponse{
		RID:   s.GetRID(),
		Users: waiters,
	}); err != nil {
		fmt.Printf("UsersInfo error: %s\n", err)
	}
	fmt.Printf("Handle UsersInfoEvent: %+v\n", s)
}

func TablesInfoEvent(e *expirenza.Expirenza, s *expirenza.TablesInfoEvent) {
	if err := e.TablesInfo(context.Background(), &expirenza.TablesInfoResponse{
		RID:    s.GetRID(),
		Tables: tables,
	}); err != nil {
		fmt.Printf("TablesInfo error: %s\n", err)
	}
	fmt.Printf("Handle TablesInfoEvent: %+v\n", s)
}

func HallPlansInfoEvent(e *expirenza.Expirenza, s *expirenza.HallPlansInfoEvent) {
	if err := e.HallPlansInfo(context.Background(), &expirenza.HallPlansInfoResponse{
		RID:       s.GetRID(),
		HallPlans: hallPlans,
	}); err != nil {
		fmt.Printf("HallPlansInfo error: %s\n", err)
	}
	fmt.Printf("Handle HallPlansInfoEvent: %+v\n", s)
}

func GetBillEvent(e *expirenza.Expirenza, s *expirenza.GetBillEvent) {
	fmt.Printf("Handle GetBillEvent: %+v\n", s)
}

func PayBillRequest(e *expirenza.Expirenza, s *expirenza.PayBillRequest) {
	fmt.Printf("Handle PayBillRequest: %+v\n", s)
}

func CreateOrderOnTableEvent(e *expirenza.Expirenza, s *expirenza.CreateOrderOnTableEvent) {
	fmt.Printf("Handle CreateOrderOnTableEvent: %+v\n", s)
}

func CheckProductsRestrictionsEvent(e *expirenza.Expirenza, s *expirenza.CheckProductsRestrictionsEvent) {
	fmt.Printf("Handle CheckProductsRestrictionsEvent: %+v\n", s)
}

func SplitOrderEvent(e *expirenza.Expirenza, s *expirenza.SplitOrderEvent) {
	fmt.Printf("Handle SplitOrderEvent: %+v\n", s)
}

func CategoriesInfoEvent(e *expirenza.Expirenza, s *expirenza.CategoriesInfoEvent) {
	fmt.Printf("Handle CategoryInfoEvent: %+v\n", s)
}

func CheckWorkEvent(e *expirenza.Expirenza, s *expirenza.CheckWorkEvent) {
	fmt.Printf("Handle CheckWorkEvent: %+v\n", s)
}

func SetSettingsEvent(e *expirenza.Expirenza, s *expirenza.SetSettingsEvent) {
	fmt.Printf("Handle SetSettingsEvent: %+v\n", s)
}

func StopListEvent(e *expirenza.Expirenza, s *expirenza.StopListEvent) {
	fmt.Printf("Handle StopListEvent: %+v\n", s)
}

func VersionInfoEvent(e *expirenza.Expirenza, s *expirenza.VersionInfoEvent) {
	fmt.Printf("Handle VersionInfoEvent: %+v\n", s)
	if err := e.VersionInfo(context.Background(), &expirenza.VersionInfoResponse{
		PartnerID:                  "This, is elon musk",
		PluginCompilationTimestamp: "2023.12.12 12:12:12",
		PluginVersion:              "v0.6.0",
		RID:                        s.GetRID(),
		RestaurantLicenseDateEnd:   "2051.05.05",
		RestaurantVersion:          "v1.0.0",
		SystemInfo: expirenza.SystemInfo{
			AskCashierForMultiCashRegisterPayment: true,
			HostTerminalName:                      "Terminal",
			IsRoundingDownOrderSum:                true,
			IsMain:                                true,
		},
		UpdateInfo: []expirenza.UpdateInfo{
			{
				Enabled:        true,
				LastRunTime:    "2023.12.12 12:12:12",
				LastTaskResult: 0,
				TaskName:       "Update",
			},
		},
	}); err != nil {
		fmt.Printf("VersionInfo error: %s\n", err)
	}
}

func MenuInfoEvent(e *expirenza.Expirenza, s *expirenza.MenuInfoEvent) {
	fmt.Printf("Handle MenuInfoEvent: %+v\n", s)
}
