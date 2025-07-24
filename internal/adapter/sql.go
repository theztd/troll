package adapter

import (
	"fmt"
	"log"
	"strings"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	//	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/theztd/troll/internal/config"
)

func getDriverNameFromDSN(dsn string) (string, error) {
	switch {
	case strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://"):
		return "postgres", nil
	case strings.HasPrefix(dsn, "mysql://"):
		return "mysql", nil
	case strings.HasPrefix(dsn, "sqlite://"):
		return "sqlite", nil
	default:
		return "", fmt.Errorf("ERR [getDriverNameFromDSN]: Unknown DSN: %s", dsn)
	}
}

func NormalizeRow(row map[string]interface{}) map[string]interface{} {
	/*
		Normalize row
		- []bytes -> string
		- time object -> rfc time
	*/
	out := make(map[string]interface{})
	for k, v := range row {
		switch val := v.(type) {
		case []byte:
			out[k] = string(val)

		case time.Time:
			out[k] = val.Format(time.RFC3339)
		default:
			out[k] = val
		}
	}
	return out
}

func RunQuery(query string) (results []map[string]interface{}, err error) {
	// log.Printf("DEBUG [RunQuery]: Query: %s\n", query)
	driver, err := getDriverNameFromDSN(config.DSN)
	if err != nil {
		log.Printf("ERR [RunQuery]: Unable to parse DSN (%s). %v", config.DSN, err)
		return nil, fmt.Errorf("unable to parse DSN (%s). %v", config.DSN, err)
	}

	db, err := sqlx.Connect(driver, config.DSN)
	if err != nil {
		log.Printf("ERR [RunQuery]: Unable to connect database. %v", err)
		return nil, fmt.Errorf("unable to connect database. %v", err)
	}
	defer db.Close()

	rows, err := db.Queryx(query)
	if err != nil {
		log.Printf("ERR [RunQuery]: Database query error. %v", err)
		return nil, fmt.Errorf("database query error. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			log.Printf("WARN [RunQuery]: Skipping row due to MapScan error: %v", err)
			continue
		}
		results = append(results, NormalizeRow(row))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed: %w", err)
	}

	return results, nil
}
