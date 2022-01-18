package db

import (
	"context"

	"github.com/amaru0601/go-rent/ent"
	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 user=postgres dbname=go-rent password=123456 sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
