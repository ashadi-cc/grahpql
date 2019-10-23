package pg

import (
	"database/sql"
	"fmt"
	"gql-ashadi/service/config"
	"gql-ashadi/service/logger"
	"sync"

	// postgres driver
	_ "github.com/lib/pq"
)

var db *sql.DB
var initDb sync.Once

//GetInstance get database instance
func GetInstance() *sql.DB {
	initDb.Do(func() {
		if db == nil {
			db, _ = connectDB()
		}

	})
	return db
}

//MustConnect try to connect to postgres and exit when error
func MustConnect() {
	var err error
	db, err = connectDB()
	if err != nil {
		logger.GetLogger().Fatal(err)
	}
}

func connectDB() (*sql.DB, error) {
	connString := connString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connString() string {
	cfg := config.GetConfig()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
}
