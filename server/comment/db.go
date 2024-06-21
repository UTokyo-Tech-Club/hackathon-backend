package comment

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS comment (
		uid VARCHAR(255) PRIMARY KEY, 
		post_uid VARCHAR(255), 
		comment TEXT, 
		commenting_user_uid VARCHAR(255), 
		create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP)`

	if _, err := mysql.CreateTable(query); err != nil {
		logger.Error(err)
		return
	}
}
