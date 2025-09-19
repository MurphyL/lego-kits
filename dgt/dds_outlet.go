package dgt

/* Dynamic Get Toolkit - 动态数据源 */

type DataSourceType string

type DataSourceAction string

const (
	RDBMS  DataSourceType = "rdbms"
	JSONL  DataSourceType = "jsonl"
	NDJSON DataSourceType = "ndjson"
	HTTP   DataSourceType = "http"
)

const (
	GET DataSourceAction = "get"
)

type DDS interface {
	Apply(action DataSourceAction)
}
