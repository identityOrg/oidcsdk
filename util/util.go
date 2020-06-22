package util

import (
	rand2 "crypto/rand"
	"crypto/rsa"
	"net/url"
	"strings"
)

func AppendUnique(slice []string, item string) []string {
	for _, s := range slice {
		if item == s {
			return slice
		}
	}
	return append(slice, item)
}

func StringInSlice(needle string, haystack []string) bool {
	for _, b := range haystack {
		if strings.ToLower(b) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func RemoveEmpty(args []string) (ret []string) {
	for _, v := range args {
		v = strings.TrimSpace(v)
		if v != "" {
			ret = append(ret, v)
		}
	}
	return
}

func GetAndRemove(values url.Values, key string) string {
	defer values.Del(key)
	return values.Get(key)
}

func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	key, err := rsa.GenerateKey(rand2.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return key, &key.PublicKey
}
