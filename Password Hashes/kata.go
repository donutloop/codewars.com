package Password_Hashes

import (
	"crypto/md5"
	"encoding/hex"
)

func PassHash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
