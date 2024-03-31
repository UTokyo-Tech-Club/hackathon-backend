package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"hackathon-backend/utils/logger"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// func Init() {
// 	dbUser := os.Getenv("DB_USER")
// 	dbPwd := os.Getenv("DB_PASS")
// 	dbName := os.Getenv("DB_NAME")
// 	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")

// 	dsn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)

// 	// Use cloudsqlconn to dial the Cloud SQL instance
// 	_, err := cloudsqlconn.NewDialer(context.Background())
// 	if err != nil {
// 		logger.Error("Could not create dialer: ", err)
// 	}

// 	// var opts []cloudsqlconn.DialOption

// 	// sqlcon.RegisterDialContext("cloudsqlconn",
// 	// 	func(ctx context.Context, addr string) (net.Conn, error) {
// 	// 		return d.Dial(ctx, instanceConnectionName, opts...)
// 	// 	})

// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		logger.Error("Could not open db: ", err)
// 	}

// 	// dsn := "root:TheoJang(30@tcp(34.146.51.218:3306)/hackathon"
// 	// err := error(nil)
// 	// db, err = sql.Open("mysql", dsn)
// 	// if err != nil {
// 	// 	logger.Error(err)
// 	// 	return
// 	// }

// 	// Check if the connection is successful
// 	err = db.Ping()
// 	if err != nil {
// 		logger.Error(err)
// 		return
// 	}

// 	logger.Info("Connected to the MySQL database successfully!")
// }

func Init() {

	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_connector.go: %s environment variable not set.", k)
		}
		return v
	}
	// Note: Saving credentials in environment variables is convenient, but not
	// secure - consider a more secure solution such as
	// Cloud Secret Manager (https://cloud.google.com/secret-manager) to help
	// keep passwords and other secrets safe.
	var (
		dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
		dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
		dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		usePrivate             = os.Getenv("PRIVATE_IP")
	)

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		logger.Error("Could not create dialer: ", err)
		return
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPwd, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		logger.Error("sql.Open: ", err)
		return
	}
	db = dbPool
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
