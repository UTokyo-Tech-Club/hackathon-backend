package mysql

import (
	"database/sql"
	"hackathon-backend/utils/logger"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	dsn := "root:TheoJang(30@tcp(34.146.51.218:3306)/hackathon"
	err := error(nil)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
		return
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("Connected to the MySQL database successfully!")
}

func Exec(query string, args ...interface{}) (sql.Result, error) {

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Prepare a statement within the transaction
	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the statement
	result, err := stmt.Exec(args...)
	if err != nil {

		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}

func CreateTable(query string) (sql.Result, error) {

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(query)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, nil
}
