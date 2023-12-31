package acquiring

import (
	"errors"
)

var (
	// ErrFailedToParsePEMBlock failed to parse PEM block containing the public key
	ErrFailedToParsePEMBlock = errors.New("failed to parse PEM block containing the public key")

	// ErrFailedToParsePublicKey failed to parse public key
	ErrFailedToParsePublicKey = errors.New("failed to parse public key")
)
