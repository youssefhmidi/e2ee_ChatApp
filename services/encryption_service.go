package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/encryption"
)

// this service should be used if in the future I want to add  bots
type GroupChatEncryption struct {
	GroupPublicKey []byte
}

func NewGCE(groupPubKey string) *GroupChatEncryption {
	return &GroupChatEncryption{
		GroupPublicKey: []byte(groupPubKey),
	}
}

/*
encrypte the signed message sent from the user into the group chat and returns it
*/
func (GCE *GroupChatEncryption) EncryptMessage(signedToken string) string {
	gcm := encryption.ParseGCMfromKey(GCE.GroupPublicKey)
	encrypted := encryption.EncrypteSignedMessage(gcm, signedToken)
	return encrypted
}

// decrypte the encrypted message and return the raw signed text
func (GCE *GroupChatEncryption) DecryptMessage(encrypted_message string) string {
	gcm := encryption.ParseGCMfromKey(GCE.GroupPublicKey)
	decrypted := encryption.DecrypteMessage(gcm, encrypted_message)
	return decrypted
}

// verifying the signature of a token (the token is encrypted so this is the chain "sign raw text -> encrypt the signed text -> send, recive ->
// decrypte -> verify signature"
func VerifySignature(signedToken string, publicKey string) (jwt.MapClaims, bool) {
	Pub, err := encryption.ParseKey([]byte(publicKey))
	if err != nil {
		panic(err)
	}
	mapMessage, IsVerified := encryption.Verify(signedToken, Pub)
	return mapMessage, IsVerified
}
