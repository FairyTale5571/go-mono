package wss

import "errors"

var (
	// ErrWSAlreadyOpen a websocket that already is open.
	ErrWSAlreadyOpen = errors.New("web socket already opened")

	// ErrWSNotFound is thrown when you attempt to use a websocket
	// that doesn't exist
	ErrWSNotFound = errors.New("no websocket connection exists")
)
