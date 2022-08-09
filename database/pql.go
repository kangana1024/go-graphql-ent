package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/kangana1024/go-graphql-ent/ent"
	_ "github.com/lib/pq"
)

var DB *ent.Client

func NewPGConnectionDB() error {
	client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	return nil
}
