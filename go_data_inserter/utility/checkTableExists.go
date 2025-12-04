package utility

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CheckTableExists(db *sql.DB, schemaType string, tableName string) bool {
	exists := false

	query := `SELECT EXISTS (
	SELECT FROM information_schema.tables
	WHERE table_schema = $1 AND table_name = $2
	)`

	db.QueryRow(query, schemaType, tableName).Scan(&exists)

	return exists
}