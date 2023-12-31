package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/fairytale5571/go-mono/public"
)

func main() {
	pub := public.NewPublic(public.Opts{
		Token:  "ugPVpCP6poHI1oXPC7k-p-......",
		Client: http.DefaultClient,
	})
	ctx := context.Background()
	var err error

	err = pub.SetWebhook(ctx, &public.SetWebhookRequest{
		WebHookURL: "https://b229-95-158-48-251.ngrok-free.app",
	})
	if err != nil {
		panic(err)
	}

	info, err := pub.GetInfoClient(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(info)

	currencies, err := pub.GetCurrencies(ctx)
	if err != nil {
		panic(err)
	}
	log.Println(currencies)

	statements, err := pub.GetStatements(ctx, &public.GetStatementsRequest{
		Account: "XoO4NZ3tAkQc_.......",
		From:    time.Now().Add(-time.Hour * 24 * 15).Unix(),
		To:      time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
	log.Println(statements)
}
