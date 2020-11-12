package webrtc

import "io"

// DataChannelMessage represents a message received from the
// data channel. IsString will be set to true if the incoming
// message is of the string type. Otherwise the message is of
// a binary type.
type DataChannelMessage struct {
	IsString bool

	// It is the caller's responsibility to close Data to release memory
	// resources.
	Data io.ReadCloser
}
