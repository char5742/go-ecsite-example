//go:build medium

package medium_test

import (
	"database/sql"
	"os"
	"testing"

	"char5742/ecsite-sample/test/share"
)

var tdb *sql.DB

func TestMain(m *testing.M) {
	// OpenDBForTest はテスト用のデータベース接続を返します。
	tdb = share.OpenDBForTest()
	os.Exit(m.Run())
}
