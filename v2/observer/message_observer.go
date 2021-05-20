package observer

type MessageObserver interface {
	Request(ID int64, method string, jsonData []byte, err error)
	Response(ID int64, method string, jsonData []byte, err error)
}

func NewIgnoreMessagesObserver() *IgnoreMessagesObserver {
	return &IgnoreMessagesObserver{}
}

type IgnoreMessagesObserver struct {}

func (observer *IgnoreMessagesObserver) Request(ID int64, method string, jsonData []byte, err error) {
	// intentionally blank
}

func (observer *IgnoreMessagesObserver) Response(ID int64, method string, jsonData []byte, err error) {
	// intentionally blank
}
