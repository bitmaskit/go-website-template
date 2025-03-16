package config

import (
	"log"
	"os"
)

type Database struct {
	Driver string
	Dsn    string
}

func (d *Database) Load() *Database {
	d.Driver = os.Getenv("DATABASE_DRIVER")
	d.Dsn = os.Getenv("DATABASE_DSN")

	if d.Driver == "" || d.Dsn == "" {
		log.Println("DATABASE LOAD ERROR: missing database configuration, defaulting to SQLite")
		d.Driver = "sqlite3"
		d.Dsn = "file::memory:?cache=shared"
	}
	return d
}
