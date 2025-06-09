package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sign 加密算法
func Sign(input, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, _ = h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
