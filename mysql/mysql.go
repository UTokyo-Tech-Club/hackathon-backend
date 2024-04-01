package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"hackathon-backend/utils/logger"
	"hackathon-backend/utils/variables"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDBServer() {

	var (
		dbUser                 = variables.MustGetenv("DB_USER")
		dbPwd                  = variables.MustGetenv("DB_PASS")
		dbName                 = variables.MustGetenv("DB_NAME")
		instanceConnectionName = variables.MustGetenv("INSTANCE_CONNECTION_NAME")
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

	DB, err = sql.Open("mysql", dbURI)
	if err != nil {
		logger.Error("sql.Open: ", err)
		return
	}

	if DB.Ping() != nil {
		logger.Error(err)
		return
	}

	logger.Info("MySQL connection successful")
}

func initDBLocal() {

	var (
		dbUser                 = variables.MustGetenv("DB_USER")
		dbPwd                  = variables.MustGetenv("DB_PASS")
		dbName                 = variables.MustGetenv("DB_NAME")
		instanceConnectionName = variables.MustGetenv("INSTANCE_CONNECTION_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPwd, instanceConnectionName, dbName)
	err := error(nil)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		logger.Error(err)
		return
	}

	// Check if the connection is successful
	if DB.Ping() != nil {
		logger.Error(err)
		return
	}

	logger.Info("Connected to MySQL")
}

func Init() {
	if os.Getenv("IS_DEPLOYED") == "true" {
		initDBServer()
	} else {
		initDBLocal()
	}
}

func Exec(query string, args ...interface{}) (sql.Result, error) {

	// Begin a transaction
	tx, err := DB.Begin()
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

	tx, err := DB.Begin()
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
