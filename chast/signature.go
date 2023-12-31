package chast

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

func (c *Chast) sign(data any) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	h := hmac.New(sha256.New, []byte(c.key))
	h.Write(bytes)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
