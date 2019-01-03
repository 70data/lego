package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MakeMD is make string md5.
func MakeMD(initString string) string {
	m := md5.New()
	m.Write([]byte(initString))
	md := m.Sum(nil)
	mdString := hex.EncodeToString(md)
	return mdString
}
