package database

import (
	"errors"
	"fmt"
	"github.com/aprilboiz/flight-management/pkg/config"
	"log"
	"os"
	"time"

	"github.com/aprilboiz/flight-management/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func GetDatabase() *gorm.DB {
	if database != nil {
		return database
	}
	db, err := initialize(zap.L())
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	return db
}

func initialize(log *zap.Logger) (*gorm.DB, error) {
	cfg := config.GetConfig()

	// Configure GORM logger
	gormLogger := logger.New(
		zap.NewStdLog(log),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(config.GetDatabaseConnectionString()), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	log.Info("Connected to database", zap.String("dsn", config.GetDatabaseConnectionString()))

	if cfg.Environment == config.EnvironmentDevelopment {
		log.Warn("Dropping all tables (Development only!)")
		if err := dropAllTables(db); err != nil {
			log.Error("Failed to drop tables", zap.Error(err))
		}
	}

	log.Info("Migrating database schema")
	if err := migrateDatabase(db); err != nil {
		log.Error("Failed to migrate database schema", zap.Error(err))
		return nil, err
	}

	database = db
	log.Info("Initialized database successfully")
	if cfg.Database.Init.RunSeed {
		log.Info("Running seed script")
		err := runSQLScript(cfg.Database.Init.SeedPath)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func dropAllTables(db *gorm.DB) error {
	tables, err := db.Migrator().GetTables()
	if err != nil {
		zap.L().Error("Error dropping tables", zap.Error(err))
	} else {
		for _, table := range tables {
			if err := db.Migrator().DropTable(table); err != nil {
				zap.L().Error("Error dropping table", zap.String("table", table), zap.Error(err))
			}
		}
	}
	return err
}

func migrateDatabase(db *gorm.DB) error {
	//err := db.SetupJoinTable(&models.Flight{}, "IntermediateStops", &models.IntermediateStop{})
	//if err != nil {
	//	return err
	//}
	return db.AutoMigrate(
		&models.Plane{},
		&models.TicketClass{},
		&models.Airport{},
		&models.Seat{},
		&models.Flight{},
		&models.IntermediateStop{},
		&models.Ticket{},
		&models.Configuration{},
		//&models.FlightTicketDetails{},
	)
}

func GetSequenceNameForTable(table string, column string) (string, error) {
	db := database
	if db == nil {
		return "", errors.New("database connection is nil")
	}

	var sequenceName string
	err := db.Raw("SELECT pg_get_serial_sequence($1, $2)", table, column).Scan(&sequenceName).Error
	if err != nil {
		zap.L().Error("Error getting sequence name for table", zap.String("table", table), zap.Error(err))
		return "", err
	}
	return sequenceName, err
}

func runSQLScript(sqlFilePath string) error {
	// Read the SQL file
	sqlBytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	// Get the database connection
	db := database
	if db == nil {
		return errors.New("database connection is nil")
	}

	// Execute the SQL script
	err = db.Exec(string(sqlBytes)).Error
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %w", err)
	}

	zap.L().Info("SQL script executed successfully")
	return nil
}

func PeekUpcomingFlightId() (uint, error) {
	db := database
	var maxId uint

	// Query the maximum ID currently in the flights table
	err := db.Raw("SELECT COALESCE(MAX(id), 0) FROM flights").Scan(&maxId).Error
	if err != nil {
		zap.L().Error("Error getting max flight ID", zap.Error(err))
		return 0, err
	}

	// Return max+1 as the next ID
	return maxId + 1, nil
}
