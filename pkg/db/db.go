package db

import (
	"fmt"
	"log"

	"bitbucket.org/Local/Scouter/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig holds the database configuration parameters
type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	Server   string
}

// ConnectDB initializes the PostgreSQL database connection
func ConnectDB() *gorm.DB {
	// Define database configuration
	config := DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "1",
		DBName:   "postgres",
		Port:     "5433",
		Server:   "ScouterDB",
	}

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.DBName, config.Port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[%s] ❌ Failed to connect to database: %v", config.Server, err)
	}

	fmt.Printf("[%s] ✅ Successfully connected to PostgreSQL!\n", config.Server)
	return DB
}

// MigrateDB creates the necessary tables in the database
func MigrateDB() {
	if DB == nil {
		log.Fatal("❌ Database connection is not initialized")
	}

	err := DB.AutoMigrate(&models.Team{}, &models.Player{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	fmt.Println("✅ Database migration completed successfully!")
}
