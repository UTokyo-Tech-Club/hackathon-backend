package mysql

import (
	"database/sql"
	"hackathon-backend/utils/logger"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dsn := "root:TheoJang(30@tcp(34.146.51.218:3306)/hackathon"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Connected to the MySQL database successfully!")
}
