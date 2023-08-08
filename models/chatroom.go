package models

import "gorm.io/gorm"

type ChatRoom struct {
	gorm.Model
	OwnerID   uint
	Name      string
	IsPublic  bool
	PublicKey string
	Members   []User    `gorm:"many2many:user_chatroom"`
	Messages  []Message `gorm:"foreignKey:ChatRoomID"`
}
