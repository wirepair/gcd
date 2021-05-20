package observer

type MessageObserver interface {
}

type IgnoreMessagesObserver struct {}

func NewIgnoreMessagesObserver() *IgnoreMessagesObserver {
	return &IgnoreMessagesObserver{}
}
