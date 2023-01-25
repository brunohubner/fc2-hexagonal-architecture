package db

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func SetUpTests() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func Close() error {
	if err := Db.Close(); err != nil {
		return err
	}

	return nil
}

func createTable(db *sql.DB) {
	table := `
		CREATE TABLE IF NOT EXISTS products (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			price TEXT REAL DEFAULT 0,
			status TEXT DEFAULT "disabled"
	);
	`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := fmt.Sprintf(`
		insert into products values(
			"%s", "Product test", 0, "disabled"
		);
	`, "9235d7aa-7854-4f50-8d4d-0fd4d3263c0a")

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}
