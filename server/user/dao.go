package user

type Dao interface {
	Register(userData UserData) error
}
