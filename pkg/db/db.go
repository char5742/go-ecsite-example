//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package db

import (
	"char5742/ecsite-sample/pkg/config"
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq" // ドライバのインポート
)

// DB interface to abstract the database operations
type TX interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Rollback() error
	Commit() error
}

type TXImpl struct {
	*sql.Tx
}

func NewTX(tx *sql.Tx) TX {
	return &TXImpl{tx}
}

type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TX, error)
	OpenMigrate(sourceUrl string) *migrate.Migrate
}

type DBImpl struct {
	*sql.DB
}

func NewDB(db *sql.DB) DB {
	return &DBImpl{db}
}

func (t *DBImpl) BeginTx(ctx context.Context, opts *sql.TxOptions) (TX, error) {
	tx, err := t.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return NewTX(tx), nil
}
func (db *DBImpl) OpenMigrate(sourceUrl string) *migrate.Migrate {
	return OpenMigrate(db.DB, sourceUrl)
}

var (
	conn     *sql.DB
	initOnce sync.Once
	initErr  error
)

// GetDBConn はスレッドセーフでシングルトンの DatabaseConnection を返します。
// 初回呼び出し時のみ接続を確立し、エラーがあれば呼び出し元に返します。
func OpenDB() (DB, error) {
	initOnce.Do(func() {
		_conn, err := open()
		if err != nil {
			initErr = fmt.Errorf("failed to open database: %w", err)
			return
		}

		// 接続確認
		if err := _conn.Ping(); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}
		conn = _conn

		// コネクションプール設定（必要に応じて値を調整）
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(5)
		conn.SetConnMaxLifetime(time.Hour)

	})
	return NewDB(conn), initErr
}

// open は config から取得した設定をもとにデータベースへ接続し、*sql.DB を返します。
func open() (*sql.DB, error) {
	cfg := config.GetConfig()
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)
	fmt.Print(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}
