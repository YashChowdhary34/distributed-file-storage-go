package p2p

import "errors"

// ErrInvalidHandshake is returned if the handshake betweek
// the local and remote node could not be estabilished
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(any) error

func NOPHandshakeFunc(any) error { return nil }