## ClickHouse SQL Parser Issue

This is just to show how the ClickHouse SQL Parser library fails to parse queries with named parameters (@ syntax).

There are two test suites that run the same queries:
- On ClickHouse (adapt the connection string to your instance)
- On the ClickHouse SQL Parser library

The same queries work in ClickHouse but fail in the parser.


```bash
$ go test -timeout 30s -run ^TestClickHouseSQLParserLibraryTest$ clickhouse-parser-test -count 1
--- FAIL: TestClickHouseSQLParserLibraryTest (0.00s)
    parser_test.go:23: error: line 0:7 unexpected token kind: @
        SELECT @field_name as result
               ^
    parser_test.go:23: error: line 0:21 expected table name or subquery, got @
        SELECT count(*) FROM @table_name
Initial commit.
                             ^
```