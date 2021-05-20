package observer

import "github.com/wirepair/gcd/v2/gcdmessage"

type MessageObserver interface {
	Request(ID int64, method string, jsonData []byte)
	Response(ID int64, method string, jsonData []byte, err error)
}

// DigResponseData returns the response data if there is any
func DigResponseData(response *gcdmessage.Message) []byte {
	if response == nil {
		return nil
	}

	return response.Data
}


func NewIgnoreMessagesObserver() *IgnoreMessagesObserver {
	return &IgnoreMessagesObserver{}
}

type IgnoreMessagesObserver struct {}

func (observer *IgnoreMessagesObserver) Request(ID int64, method string, jsonData []byte) {
	// intentionally blank
}

func (observer *IgnoreMessagesObserver) Response(ID int64, method string, jsonData []byte, err error) {
	// intentionally blank
}
