package main

import (
	"char5742/ecsite-sample/internal/share/infra/db"
	"os"
)

func main() {
	conn, err := db.OpenDB()
	if err != nil {
		panic(err)
	}
	args := os.Args[1]
	m := db.OpenMigrate(conn.Conn, "file://db/migrations")
	defer m.Close()
	if args == "up" {
		if err := m.Up(); err != nil {
			panic(err)
		}
	} else if args == "down" {
		if err := m.Down(); err != nil {
			panic(err)
		}
	} else {
		panic("invalid argument error")
	}
	os.Exit(0)
}
