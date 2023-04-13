package tpcc

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestPrep(t *testing.T) {
	dsn := fmt.Sprintf("root:@tcp(127.0.0.1:4000)/test?multiStatements=true")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	conn.QueryRowContext(context.Background(), "select * from t where a < 1")
}
