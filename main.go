package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

//Create connection to database as per given database file path
func CreateConnection() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db2, err := gorm.Open(sqlite.Open("springct.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db = db2

	db.AutoMigrate(&Course{})
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&CourseStudent{})

}

func main() {
	// Create Server and Route Handlers
	CreateConnection()

	handleWeb()
	//}()

	// Graceful Shutdown
	//waitForShutdown(srv)
}
