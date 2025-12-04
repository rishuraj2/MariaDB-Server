package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go_data_inserter/utility"
)

const (
	PORT = ":8081"
	schemaType = "public"
	tableName = "batch_inputs"
)

func main() {
	db, err := utility.ConnectDb();
	defer db.Close()
	if err != nil {
		log.Fatalf("Error in connection to database. Check connection string. err: %s", err)
	}

	if err := utility.CheckDbConnection(db); err != nil {
		log.Fatalf("Failed to connect to database: ", err)
		return
	}

	if tableExists := utility.CheckTableExists(db, schemaType, tableName); tableExists {
		log.Printf("Table found and ready!")
	} else {
		log.Fatalf("Failed to find table %s! Check 'Schema Type' and 'Table Name'", tableName)
	}
}