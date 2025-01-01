package share

import (
	"char5742/ecsite-sample/internal/share/infra/db"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func OpenDBForTest() *sql.DB {

	ctx := context.Background()
	pgSql, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres:17",
			Env: map[string]string{
				"POSTGRES_USER":     "test",
				"POSTGRES_PASSWORD": "test",
				"POSTGRES_DB":       "test",
				"TZ":                "UTC",
			},
			ExposedPorts: []string{"5432/tcp"},
			WaitingFor:   wait.ForListeningPort("5432/tcp"),
		},
		Started: true,
	})

	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	host, err := pgSql.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get host: %v", err)
	}
	port, err := pgSql.MappedPort(ctx, "5432/tcp")
	if err != nil {
		log.Fatalf("failed to get externally mapped port: %s", err)
	}

	tdb, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port.Port(),
			"test",
			"test",
			"test",
		),
	)

	if err != nil {
		log.Fatalf("failed to open a connection to the database: %s", err)
	}

	if err := tdb.Ping(); err != nil {
		log.Fatalf("failed to verify a connection to the database: %s", err)
	}

	m := db.OpenMigrate(tdb, "file://../../../db/migrations")
	m.Up()
	return tdb
}

func ExecSQLFile(t *testing.T, path string, dbConn *sql.DB) {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read sql file: %v", err)
	}
	if _, err := dbConn.Exec(string(sqlBytes)); err != nil {
		t.Fatalf("failed to exec sql file: %v", err)
	}
}
