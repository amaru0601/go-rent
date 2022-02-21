package db

import (
	"context"
	"log"

	"github.com/amaru0601/go-rent/ent"
	"github.com/amaru0601/go-rent/ent/migrate"
	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 dbname=go-rent user=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
		return nil, err
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
		return nil, err
	}

	return client, nil
}
