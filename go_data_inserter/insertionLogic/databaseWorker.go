package insertionlogic

import (
	"log"

	"database/sql"
	"go_data_inserter/models"
)

func DatabaseWorker(db *sql.DB, dataChan chan models.DataToBeInserted) {
	for data := range dataChan {
		insertBatchData(db, data)
	}
}

func insertBatchData(db *sql.DB, data models.DataToBeInserted) {
	_, err := db.Exec(`INSERT INTO batch_inputs(start_time, raw_material_name, raw_material_quantity, temperature, end_time)
	VALUES($1, $2, $3, $4, $5)
	`, data.StartTime, data.RawMaterialName, data.RawMaterialQuantity, data.Temperature, data.EndTime)

	if err != nil {
		log.Fatal("Error inserting data: ", err)
	} else {
		log.Printf("Inserted Batch Data: raw_material_name: %s, raw_material_quantity: %d | %s -> %s", data.RawMaterialName, data.RawMaterialQuantity, data.StartTime.Format("15:04:05"), data.EndTime.Format("15:04:05"))
	}
}