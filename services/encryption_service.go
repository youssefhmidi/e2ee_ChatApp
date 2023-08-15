package services

type GroupChatEncryption struct {
	GroupPublicKey []byte
}

/*
encrypte the message the signed message sent from the user into the group chat and returns it
*/
func (GCE *GroupChatEncryption) EncryptMessage(signedToken string) string {
	return ""
}

func (GCE *GroupChatEncryption) DecryptMessage(encrypted_message string) {

}

func (GCE *GroupChatEncryption) GetSender(signedToken string, MembersPublicKey []string) {

}
