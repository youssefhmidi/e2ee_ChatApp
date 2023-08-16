package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// this internal package is mostly used for public group chat

// Creates a symetric key for the groupchat
func CreateSymetricKey() (cipher.AEAD, string) {
	rawKey := make([]byte, 32)
	rand.Read(rawKey)
	Key := base64.StdEncoding.EncodeToString(rawKey)[:32]
	aes, err := aes.NewCipher([]byte(Key))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}
	return gcm, Key
}

// used for converting a base64 string to a gcm
func ParseGCMfromKey(key []byte) cipher.AEAD {
	Aes, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(Aes)
	if err != nil {
		panic(err)
	}
	return gcm
}

// encrypte the signed message that the user have sent
func EncrypteSignedMessage(gcm cipher.AEAD, signedMessage string) string {
	nonce := make([]byte, gcm.NonceSize())
	_, err := rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	return string(gcm.Seal(nonce, nonce, []byte(signedMessage), nil))
}

// gets the signed string from the encrypted signed string and returns it
func DecrypteMessage(gcm cipher.AEAD, cipherText string) string {
	nonce := gcm.NonceSize()
	// ceperating the nounce from the cipher text
	Encryptednonce, EncryptedText := cipherText[:nonce], cipherText[nonce:]
	decryptedText, err := gcm.Open(nil, []byte(Encryptednonce), []byte(EncryptedText), nil)

	if err != nil {
		panic(err)
	}
	return string(decryptedText)
}

// Verify function verifies who send the message and returns the raw text
// and a boolean , if an error happen it will panic and stop the program
func Verify(decryptedText string, publicKey *rsa.PublicKey) (jwt.MapClaims, bool) {
	tok, err := jwt.Parse(decryptedText, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("not a valid signing method")
		}
		return publicKey, nil
	})
	if err != nil {
		return jwt.MapClaims{}, false
	}

	return tok.Claims.(jwt.MapClaims), true
}

// this parses the public keys which are stored in the database from []byte into *rsa.PublicKey
func ParseKey(pem []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(pem)
}
