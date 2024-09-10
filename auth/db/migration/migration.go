package migration

import (
	"biodata/db"
	model "biodata/model/entity"
	"fmt"
	"log"
)

func MigrationHandler() {
	err := db.DbInit().AutoMigrate(&model.DataIdentitas{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}