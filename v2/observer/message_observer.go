package observer

var Observe MessageObserver = &IgnoreAllMessagesObserver{}

type MessageObserver interface {
	Request(ID int64, method string, jsonData []byte, err error)
	Response(ID int64, method string, jsonData []byte, err error)
}

type IgnoreAllMessagesObserver struct {}

func (*IgnoreAllMessagesObserver) Request(ID int64, method string, jsonData []byte, err error) {
	// intentionally blank
}

func (*IgnoreAllMessagesObserver) Response(ID int64, method string, jsonData []byte, err error) {
	// intentionally blank
}
