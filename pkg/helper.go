package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
)

func GenerateRandomKey(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)[:length]
}

func GenerateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func NormalizeString(s string) string {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")
	return s
}
