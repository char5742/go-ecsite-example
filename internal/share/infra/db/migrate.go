package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func OpenMigrate(
	db *sql.DB,
	sourceUrl string,
) *migrate.Migrate {
	// マイグレーション処理
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create a new postgres driver: %s", err)
	}
	// マイグレーションのパスは実行ファイルからの相対パスで指定する
	m, err := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed to create a new migrate instance: %s", err)
	}
	return m
}
