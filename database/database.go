package database

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"setup/ent"
)

type Idatabase interface {
	main()
}

var client *ent.Client
var err error

type MySQL struct{}
type SQLite struct{}
type PostgreSQL struct{}

func Init() {
	if err := entc.Generate("./ent/schema", &gen.Config{}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
	var database MySQL //If you want to use SQLite or PostgreSQL, use instead of MySQL.
	Idatabase.main(&database)
}
func DBManager() *ent.Client {
	return client
}
