package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

//type DbInstance struct {
//	sqlDB *sql.DB
//	DB    *gorm.DB
//}

var (
	globalDB *gorm.DB
	dbOnce   sync.Once
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=wetao_db password=REWQ_7AD83439wEqwR dbname=wetao_db port=5477 sslmode=disable TimeZone=Europe/Moscow"
	var err error

	dbOnce.Do(func() {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			QueryFields: true,
			PrepareStmt: true,
			//DryRun:      true,
			Logger: logger.Default.LogMode(logger.Warn),
		})
		if err != nil {
			panic(err)
		}

		globalDB = db
	})

	return globalDB, err
}

func GetDB() *gorm.DB {
	return globalDB
}
