package database

import (
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
	
	var database MySQL //If you want to use SQLite or PostgreSQL, use instead of MySQL.
	Idatabase.main(&database)
}
func DBManager() *ent.Client {
	return client
}
