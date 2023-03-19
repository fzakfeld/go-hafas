package hafas

import (
	"crypto/md5"
	"encoding/hex"
)

func (c *hafasClient) createChecksum(requestData []byte) string {
	salt, err := hex.DecodeString(c.salt)
	if err != nil {
		panic(err)
	}

	data := append(requestData, salt...)
	hash := md5.Sum(data)

	return hex.EncodeToString(hash[:])
}
