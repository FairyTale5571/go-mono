package main

import (
	"context"
	"fmt"

	"github.com/FairyTale5571/go-mono/expirenza"
)

func main() {
	client := expirenza.NewExpirenzaClient(expirenza.Settings{
		RestoID:   "resto_id_here",
		SecretKey: "key_here",
		Debug:     true,
	})

	ctx := context.Background()

	client.AddHandler(func(e *expirenza.Expirenza, s *expirenza.UsersInfoRequest) {
		if err := e.UsersInfo(context.Background(), &expirenza.UsersInfoResponse{
			RID:   s.RID,
			Users: waiters,
		}); err != nil {
			fmt.Printf("UsersInfo error: %s\n", err)
		}
		fmt.Printf("Handle UsersInfoRequest: %+v\n", s)
	})
	client.AddHandler(func(e *expirenza.Expirenza, s *expirenza.TablesInfoRequest) {
		if err := e.TablesInfo(context.Background(), &expirenza.TablesInfoResponse{
			RID:    s.RID,
			Tables: tables,
		}); err != nil {
			fmt.Printf("TablesInfo error: %s\n", err)
		}
		fmt.Printf("Handle TablesInfoRequest: %+v\n", s)
	})
	client.AddHandler(func(e *expirenza.Expirenza, s *expirenza.CreateOrderOnTableRequest) {
		if err := e.HallPlansInfo(context.Background(), &expirenza.HallPlansInfoResponse{
			RID:       s.RID,
			HallPlans: hallPlans,
		}); err != nil {
			fmt.Printf("HallPlansInfo error: %s\n", err)
		}
		fmt.Printf("Handle CreateOrderOnTableRequest: %+v\n", s)
	})

	if err := client.Open(ctx); err != nil {
		panic(err)
	}

	select {}
}
