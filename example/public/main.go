package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FairyTale5571/go-mono/open-api"
)

func main() {
	pub := open_api.NewPublic(open_api.Opts{
		Token:  "ugPVpCP6poHI1oXPC7k-p-......",
		Client: http.DefaultClient,
	})
	ctx := context.Background()
	var err error

	err = pub.SetWebhook(ctx, &open_api.SetWebhookRequest{
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

	statements, err := pub.GetStatements(ctx, &open_api.GetStatementsRequest{
		Account: "XoO4NZ3tAkQc_.......",
		From:    time.Now().Add(-time.Hour * 24 * 15).Unix(),
		To:      time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
	log.Println(statements)
}
