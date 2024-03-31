package user

import (
	"hackathon-backend/mysql"
	"hackathon-backend/utils/logger"
)

type Dao interface {
	Register(d UserData) error
	Edit(d UserData) error
	GetProfileContent(d UserData) (UserData, error)
}

type dao struct{}

func NewDao() Dao {
	return &dao{}
}

func (dao *dao) Register(d UserData) error {
	query := "INSERT IGNORE INTO user (uid, username, email, photo_url) VALUES (?, ?, ?, ?)"
	if _, err := mysql.Exec(query, d.UID, d.Username, d.Email, d.PhotoURL); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) Edit(d UserData) error {
	query := `UPDATE user 
			SET username = ?, photo_url = ?, profile_content = ?
			WHERE uid = ?`
	if _, err := mysql.Exec(query, d.Username, d.PhotoURL, d.ProfileContent, d.UID); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (dao *dao) GetProfileContent(d UserData) (UserData, error) {
	query := "SELECT profile_content FROM user WHERE uid = ?"
	stmt, err := mysql.DB.Prepare(query)
	if err != nil {
		logger.Error(err)
		return UserData{}, err
	}
	defer stmt.Close()

	var profileContent []byte
	err = stmt.QueryRow(d.UID).Scan(&profileContent)
	if err != nil {
		logger.Error(err)
		return UserData{}, err
	}

	d.ProfileContent = profileContent
	return d, nil
}
