package oauth2_oidc_sdk

import (
	"math/rand"
	"time"
)

const ValidIdChars = "1234567890abcdefghijklmnopqrstuvwxyz"

func RandomIdString(length int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = ValidIdChars[seededRand.Intn(len(ValidIdChars))]
	}
	return string(b)
}
