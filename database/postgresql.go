package database

import (
	"context"
	"log"
	"setup/env"

	"setup/ent"
	_ "github.com/lib/pq"

)

func (d PostgreSQL) main() {
	var dns = env.Getenv(env.DB_USERNAME) + ":" + env.Getenv(env.DB_PASSWORD) + "@tcp(" + env.Getenv(env.DB_HOST) + ":" + env.Getenv(env.DB_PORT) + ")/" + env.Getenv(env.DB_DATABASE) + "?charset=utf8mb4&parseTime=True&loc=Local"

	client, err := ent.Open("postgres", dns)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
