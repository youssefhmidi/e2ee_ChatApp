package encryption

// an instance that uses the public key to encrypt messages
type EncryptionInsterface interface {
	// encrypte a message content and return the encrpyted message using the public key that the struct who implement this interface provide
	DecrypteMessage(message string) (string, error)
	// Note : the encryption part happens in the client side of the chat app
	// also the Privet jey it stored localy in the Client device in a file
}
