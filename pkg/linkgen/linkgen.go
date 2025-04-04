package linkgen

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const ALPHABET = "12345678910ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateShortLink(siteUrl string) (string, error) {
	h := md5.New()
	h.Write([]byte(siteUrl))
	hexCode := hex.EncodeToString(h.Sum(nil)[:7])
	decCode, err := strconv.ParseInt(hexCode, 16, 64)
	if err != nil {
		return "", err
	}
	crypted := ""
	for decCode != 0 {
		crypted += string(ALPHABET[decCode%int64(len(ALPHABET))])
		decCode /= int64(len(ALPHABET))

	}

	return crypted, nil
}
