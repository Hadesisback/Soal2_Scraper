package database

import (
    "fmt"
    "log"
    "scraper/models"
    "github.com/spf13/viper"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    // Load the config file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")

    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }

    // Retrieve database configuration
    host := viper.GetString("database.host")
    user := viper.GetString("database.user")
    password := viper.GetString("database.password")
    dbname := viper.GetString("database.dbname")
    port := viper.GetInt("database.port")
    sslmode := viper.GetString("database.sslmode")

    // Build the DSN
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbname, port, sslmode)

    // Connect to the database
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrate the Gainer model
    DB.AutoMigrate(&models.Gainer{})
}
