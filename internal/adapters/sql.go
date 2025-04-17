package adapter

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
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
		return "", fmt.Errorf("nepodporovaný DSN: %s", dsn)
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
		log.Fatal(err)
	}

	db, err := sqlx.Connect(driver, config.DSN)
	if err != nil {
		log.Fatalf("ERR [RunQuery]: Chyba připojení k DB: %v", err)
	}
	defer db.Close()

	rows, err := db.Queryx("SELECT * from employee;")
	if err != nil {
		log.Printf("ERR [RunQuery]: Chyba při dotazu: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			log.Printf("ERR [RunQuery]: Chyba MapScan: %v", err)
			continue
		}
		normalizedRow := NormalizeRow(row)
		fmt.Printf("%#v\n", normalizedRow)
		results = append(results, row)
	}

	return results, nil
}
