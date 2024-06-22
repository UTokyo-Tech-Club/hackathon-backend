package tweet

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS tweet (
		uid VARCHAR(255) PRIMARY KEY, 
		owner_uid VARCHAR(255), 
		content TEXT, 
		image_url VARCHAR(255),
		create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`

	if _, err := mysql.CreateTable(query); err != nil {
		logger.Error(err)
		return
	}

	// Add new column if it does n`ot exist
	addColumnQuery := `ALTER TABLE tweet ADD COLUMN IF NOT EXISTS image_url VARCHAR(255)`
	if _, err := mysql.Exec(addColumnQuery); err != nil {
		logger.Error(err)
		return
	}
}
