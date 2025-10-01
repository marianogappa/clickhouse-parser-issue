package main

import (
	"context"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func TestClickHouseNamedParametersSupportTest(t *testing.T) {
	// Connect to ClickHouse
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "user",
			Password: "pass",
		},
	})

	ctx := context.Background()

	// Create a test table
	createTable := "CREATE TABLE IF NOT EXISTS test_table (id UInt32, name String) ENGINE = Memory"
	err = conn.Exec(ctx, createTable)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}

	// Clean up table at the end
	defer func() {
		dropTable := "DROP TABLE IF EXISTS test_table"
		conn.Exec(ctx, dropTable)
	}()

	// Test 1: Named parameter for field name
	query1 := "SELECT @field_name as result"
	rows, err := conn.Query(ctx, query1, clickhouse.Named("field_name", "hello"))
	if err != nil {
		t.Errorf("Query with field parameter failed: %v", err)
	} else {
		rows.Close()
	}

	// Test 2: Named parameter for table name after FROM clause
	query2 := "SELECT count(*) FROM @table_name"
	rows, err = conn.Query(ctx, query2, clickhouse.Named("table_name", "test_table"))
	if err != nil {
		t.Errorf("Query with table parameter failed: %v", err)
	} else {
		rows.Close()
	}
}
