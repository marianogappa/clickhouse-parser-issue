package main

import (
	"testing"

	clickhouse "github.com/AfterShip/clickhouse-sql-parser/parser"
)

func TestClickHouseSQLParserLibraryTest(t *testing.T) {
	// ClickHouse SQL Parser does NOT support named parameters (@ syntax)
	// These are the exact same queries that work in ClickHouse but fail in the parser
	queries := []string{
		"SELECT @field_name as result",
		"SELECT count(*) FROM @table_name",
	}

	for _, query := range queries {
		parser := clickhouse.NewParser(query)
		_, err := parser.ParseStmts()
		if err == nil {
			t.Errorf("Expected parser to fail on query with named parameters: %s", query)
		} else {
			t.Logf("error: %v", err)
			t.Fail() // Just to show the error
		}
	}
}
