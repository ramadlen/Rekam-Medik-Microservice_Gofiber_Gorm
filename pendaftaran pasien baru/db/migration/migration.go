package migration

import (
	"fmt"
	"log"
	"pendaftaran_pasien_baru/db"
	model "pendaftaran_pasien_baru/model/entity"
)

func MigrationHandler() {
	err := db.DbInit().AutoMigrate(&model.DataIdentitas{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}