package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/souvik150/food-delivery/config"
)

var DB *gorm.DB

func ConnectDB(config *config.Config){
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s TimeZone=%s",
			config.DBHost,
			config.DBUserName,
			config.DBUserPassword,
			config.DBName,
			config.DBPort,
			"verify-ca",
			"ca.pem",
			"Asia/Shanghai",
		)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the Database! \n", err.Error())
        os.Exit(1)
    }

    DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
    
    DB.Logger = logger.Default.LogMode(logger.Info)

    log.Println("🚀 Connected Successfully to the Database")
}

// Config struct for database configuration
type Config struct {
    DBHost         string
    DBUserName     string
    DBUserPassword string
    DBName         string
    DBPort         string
}
