package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PublicKey      string
	Name           string
	NickName       string
	Messages       []Message  `gorm:"foreignKey:UserID"`
	ChatRoomsOwned []ChatRoom `gorm:"foreignKey:OwnerID"`
	ChatRooms      []ChatRoom `gorm:"many2many:user_chatroom"`
}
