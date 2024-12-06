//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/yawatamikiya/test2/v4/database/clickhouse"
)
