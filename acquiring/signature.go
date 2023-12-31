package acquiring

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

type InvoiceWebhook InvoiceStatusResponse

func (a *Acquiring) VerificationWebhook(pubKey, xSign string, bodyBytes []byte) (bool, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return false, err
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(xSign)
	if err != nil {
		return false, err
	}

	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return false, ErrFailedToParsePEMBlock
	}

	genericPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	pubKeyParsed, ok := genericPubKey.(*ecdsa.PublicKey)
	if !ok {
		return false, ErrFailedToParsePublicKey
	}

	hash := sha256.Sum256(bodyBytes)
	return ecdsa.VerifyASN1(pubKeyParsed, hash[:], signatureBytes), nil
}
