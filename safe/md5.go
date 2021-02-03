package safe

import (
	"crypto/md5"
	"encoding/hex"
)

func Makemd5(initString string) string {
	m := md5.New()
	m.Write([]byte(initString))
	md := m.Sum(nil)
	mdString := hex.EncodeToString(md)
	return mdString
}
