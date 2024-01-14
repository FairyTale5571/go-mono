package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FairyTale5571/go-mono/open-api/personal"
)

const (
	webhookURL = "https://domain.com"
	accountID  = "XoO4NZ3tAkQc_......."
)

func main() {
	pub := personal.NewPersonalClient(personal.Opts{
		Token:  "ugPVpCP6poHI1oXPC7k-p-......",
		Client: http.DefaultClient,
	})
	ctx := context.Background()
	var err error

	err = pub.SetWebhook(ctx, &personal.SetWebhookRequest{
		WebHookURL: webhookURL,
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

	statements, err := pub.GetStatements(ctx, &personal.GetStatementsRequest{
		Account: accountID,
		From:    time.Now().Add(-time.Hour * 24 * 15).Unix(),
		To:      time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
	log.Println(statements)
}
