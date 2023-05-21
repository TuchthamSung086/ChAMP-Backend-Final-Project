package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// postgres://hkcvsced:oEHB4yXtI4J7cWveGE4ctA_da2iSrWgr@arjuna.db.elephantsql.com/hkcvsced
	var err error
	dsn := "host=arjuna.db.elephantsql.com user=hkcvsced password=oEHB4yXtI4J7cWveGE4ctA_da2iSrWgr dbname=hkcvsced port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
