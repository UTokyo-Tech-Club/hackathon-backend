package tweet

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS tweet (
		uid VARCHAR(255) PRIMARY KEY, 
		user_uid VARCHAR(255), 
		content TEXT, 
		create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`

	if _, err := mysql.CreateTable(query); err != nil {
		logger.Fatal(err)
		return
	}
}
