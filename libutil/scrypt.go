package libutil

import (
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
)

func Scrypt(password string, salt string) string {
	scryptHex, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	scryptString := hex.EncodeToString(scryptHex)
	return scryptString
}
