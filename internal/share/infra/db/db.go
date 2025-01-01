package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"char5742/ecsite-sample/pkg/config"

	_ "github.com/lib/pq" // ドライバのインポート
)

var (
	db       *sql.DB
	initOnce sync.Once
	initErr  error
)

// GetDBConn はスレッドセーフでシングルトンの DatabaseConnection を返します。
// 初回呼び出し時のみ接続を確立し、エラーがあれば呼び出し元に返します。
func OpenDB() (*sql.DB, error) {
	initOnce.Do(func() {
		db, err := open()
		if err != nil {
			initErr = fmt.Errorf("failed to open database: %w", err)
			return
		}

		// 接続確認
		if err := db.Ping(); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}

		// コネクションプール設定（必要に応じて値を調整）
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(time.Hour)

	})

	return db, initErr
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

// func (c *DatabaseConnection) Rollback(tx *sql.Tx) error {
// 	return tx.Rollback()
// }

// func (c *DatabaseConnection) Commit(tx *sql.Tx) error {
// 	return tx.Commit()
// }

// func (c *DatabaseConnection) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
// 	return c.Conn.PrepareContext(ctx, query)
// }

// func (c *DatabaseConnection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
// 	return c.Conn.ExecContext(ctx, query, args...)
// }

// func (c *DatabaseConnection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
// 	return c.Conn.QueryContext(ctx, query, args...)
// }

// func (c *DatabaseConnection) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
// 	return c.Conn.QueryRowContext(ctx, query, args...)
// }

// func (c *DatabaseConnection) SetMaxOpenConns(n int) {
// 	c.Conn.SetMaxOpenConns(n)
// }

// func (c *DatabaseConnection) SetMaxIdleConns(n int) {
// 	c.Conn.SetMaxIdleConns(n)
// }

// func (c *DatabaseConnection) SetConnMaxLifetime(d time.Duration) {
// 	c.Conn.SetConnMaxLifetime(d)
// }

// func (c *DatabaseConnection) Ping() error {
// 	return c.Conn.Ping()
// }
