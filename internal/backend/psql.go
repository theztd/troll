package backend

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	txtTemplate "text/template"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // PG driver
)

/*
	Implementace registry/plugin strategie


	Toto je jen obalka nad dostupnymi pluginy (backendy),
	ktera umoznije dalsim pluginum se zaregistrovat
	a byt k dispozici skrz stejne misto pro zbytek aplikace
	jako factory (dejme tomu).
*/

type PSQL struct {
	db *sql.DB
}

type Table struct {
	Cols []string
	Rows [][]any
}

func normalize(v any, ct *sql.ColumnType) any {
	if v == nil {
		return nil
	}
	switch x := v.(type) {
	case []byte: // lib/pq často vrací []byte
		s := string(x)
		t := strings.ToUpper(ct.DatabaseTypeName())
		// JSON/JSONB → objekt/array
		if t == "JSON" || t == "JSONB" {
			var j any
			if json.Unmarshal(x, &j) == nil {
				return j
			}
		}
		// čísla dle typu
		switch t {
		case "INT2", "INT4", "INT8", "SMALLINT", "INTEGER", "BIGINT":
			if i, err := strconv.ParseInt(s, 10, 64); err == nil {
				return i
			}
		case "FLOAT4", "FLOAT8", "DOUBLE", "REAL", "NUMERIC", "DECIMAL":
			if f, err := strconv.ParseFloat(s, 64); err == nil {
				return f
			}
		}
		return s
	case time.Time:
		return x.UTC().Format(time.RFC3339Nano)
	case bool, int64, float64, string, json.Number, map[string]any, []any:
		return x
	default:
		return fmt.Sprint(x)
	}
}

func (pg *PSQL) Run(query string) ([]map[string]any, error) {
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colTypes, _ := rows.ColumnTypes()

	out := make([]map[string]any, 0, 64)
	for rows.Next() {
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}
		row := make(map[string]any, len(cols))
		for i, c := range cols {
			row[c] = normalize(vals[i], colTypes[i])
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

func (pg *PSQL) RunAndRenderTpl(query string, template string) ([]byte, error) {
	table, err := pg.Run(query)
	if err != nil {
		return nil, err
	}

	tmpFunc := txtTemplate.FuncMap{
		"json": func(v any) (string, error) {
			b, err := json.Marshal(v)
			return string(b), err
		},
	}

	var buf bytes.Buffer
	tmpl := txtTemplate.Must(txtTemplate.New("response").Funcs(tmpFunc).Parse(template))
	_ = tmpl.Execute(&buf, table)

	return buf.Bytes(), nil
}

func (pg *PSQL) Close() error {
	return pg.db.Close()
}

func NewPSQL(dsn string) (*PSQL, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PSQL{db: db}, nil
}
