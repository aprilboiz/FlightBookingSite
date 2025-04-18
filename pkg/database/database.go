package database

import (
	"errors"
	"strconv"
	"time"

	"github.com/aprilboiz/flight-management/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
	Timezone string
}

func GetDatabase(log *zap.Logger) *gorm.DB {
	if database != nil {
		return database
	}
	db, err := initialize(log)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}
	return db
}

func (config *Config) dsn() string {
	return "host=" + config.Host +
		" port=" + strconv.Itoa(config.Port) +
		" user=" + config.Username +
		" password=" + config.Password +
		" dbname=" + config.Dbname +
		" TimeZone=" + config.Timezone
}

func initialize(log *zap.Logger) (*gorm.DB, error) {
	config := Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Dbname:   "flight_management",
		Timezone: "Asia/Ho_Chi_Minh",
	}

	// Configure GORM logger
	gormLogger := logger.New(
		zap.NewStdLog(log), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(config.dsn()), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	log.Info("Connected to database", zap.String("dsn", config.dsn()))

	// Optional: Drop tables on startup (useful for development)
	// Comment out for production
	log.Warn("Dropping all tables (Development only!)")
	if err := dropAllTables(db); err != nil {
		log.Error("Failed to drop tables", zap.Error(err))
		// Decide if you want to return error or continue
	}

	log.Info("Migrating database schema")
	if err := migrateDatabase(db); err != nil {
		log.Error("Failed to migrate database schema", zap.Error(err))
		return nil, err
	}

	log.Info("Initialized database successfully")
	database = db
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
	err := db.SetupJoinTable(&models.Flight{}, "IntermediateStops", &models.IntermediateStop{})
	if err != nil {
		return err
	}
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

func GetNextValSequence(table string, column string) (uint, error) {
	var nextVal uint
	sequence, err := GetSequenceNameForTable(table, column)
	if err != nil {
		zap.L().Error("Error getting sequence name for table", zap.String("table", table), zap.Error(err))
		return 0, err
	}
	db := database
	err = db.Raw("SELECT nextval($1)", sequence).Scan(&nextVal).Error
	if err != nil {
		zap.L().Error("Error getting next sequence value for table", zap.String("table", table), zap.Error(err))
		return 0, err
	}
	return nextVal, nil
}
