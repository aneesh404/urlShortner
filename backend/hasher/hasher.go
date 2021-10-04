package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

func GenerateShortLink(initialLink string, userId string) string {
	// append userid to make same url from different users unique
	urlHashBytes := sha256Of(initialLink + userId) // get the unique hash

	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64() //new number generated from hash

	finalString := base64Encoded([]byte(fmt.Sprintf("%d", generatedNumber))) //encode it into base58

	return finalString[:8] // send first 8 characters of hash
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base64Encoded(bytes []byte) string {
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return string(encoded)
}
