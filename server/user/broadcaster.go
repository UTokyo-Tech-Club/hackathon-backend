package user

type Broadcaster interface {
	Follow(targetUID string, userToFollowUID string) error
}

type broadcaster struct{}

func NewBroadcaster() Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Follow(targetUID string, userToFollowUID string) error {
	return nil
}
