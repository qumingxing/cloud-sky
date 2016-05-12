package common

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

func GetGuid() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(Md5(base64.URLEncoding.EncodeToString(b), ""))
}
