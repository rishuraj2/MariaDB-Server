package main

import (
	"go_data_inserter/insertionlogic"
	"go_data_inserter/models"
	"go_data_inserter/utility"
	"log"
	"time"
)

var (
	schemaType = "public"
	tableName  = "batch_inputs"
	timeKeeper = time.Now()
	sequence   = 0
)

func main() {
	db, err := utility.ConnectDb()
	if err != nil {
		log.Fatalf("Error in connection to database. Check connection string. err: %s", err)
		return
	}
	defer db.Close()

	if err := utility.CheckDbConnection(db); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
		return
	}

	if tableExists := utility.CheckTableExists(db, schemaType, tableName); tableExists {
		log.Printf("Table found and ready!")
	} else {
		log.Fatalf("Failed to find table %s! Check 'Schema Type' and 'Table Name'", tableName)
	}

	dataChan := make(chan models.DataToBeInserted, 100)
	go insertionlogic.DatabaseWorker(db, dataChan)

	for {
		dataChan <- insertionlogic.GenerateData(&sequence, &timeKeeper)
		time.Sleep(1 * time.Second)
	}

}
