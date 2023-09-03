package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB, error) {
	return sql.Open(config_driver_name, config_data_source_name)
}

func execQuery(q string) {
	if config_debug {
		printPink(q + "\n")
	}
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Query(q)
	if err != nil {
		panic(err)
	}
}

func deleteTable(table string) {
	deleteSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", table)

	if config_debug {
		printPink(deleteSQL + "\n")
	}

	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(deleteSQL)
	if err != nil {
		panic(err)
	}
}

func showTable(table string) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	for _, column := range columns {
		printYellow(column + "	")
	}
	fmt.Println()

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Print(value + "	")
		}
		fmt.Println()
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

}
