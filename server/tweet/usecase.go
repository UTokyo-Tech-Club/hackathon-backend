package tweet

type Usecase interface {
	Post(userUID string, data []byte) (string, error)
}

type usecase struct {
	dao Dao
}

// func (u *usecase) Post(userUID string, data []byte) (string, error) {
// 	// Retrieve user data

// }
