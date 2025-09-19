package rdbms

import (
	"database/sql"
	"encoding/json"
)

type Request struct {
	DataSource *sql.DB
	SQL        string
	Args       []any
}

func (r Request) Exec() (string, error) {
	var err error
	var ret string
	if nil != r.DataSource && len(r.SQL) > 0 {
		rows, err := r.DataSource.Query(r.SQL, r.Args...)
		if err == nil && nil != rows {
			vars := ConvertDatabaseSqlRows(rows)
			data, _ := json.Marshal(vars)
			ret = string(data)
		}
	}
	return ret, err
}

type DatabaseSqlRow map[string]any

func ConvertDatabaseSqlRows(rows *sql.Rows, mappers ...func(DatabaseSqlRow)) []DatabaseSqlRow {
	columns, _ := rows.Columns()
	records := make([]DatabaseSqlRow, 0)
	for rows.Next() {
		pointers := make([]any, len(columns))
		values := make([]any, len(columns))
		for i := 0; i < len(columns); i++ {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		record := make(DatabaseSqlRow)
		for i, fieldName := range columns {
			record[fieldName] = values[i]
		}
		for _, mapper := range mappers {
			mapper(record)
		}
		records = append(records, record)
	}
	return records
}
