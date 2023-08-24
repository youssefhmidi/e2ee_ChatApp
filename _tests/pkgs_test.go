package tests

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/auth"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/encryption"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	testdata "github.com/youssefhmidi/E2E_encryptedConnection/_tests/test_data"
)

// testing all the packages in the _internals directory

func TestJWT(t *testing.T) {
	data := testdata.GetData()

	accessToken, err := auth.CreateAccessToken(data.Users[0], data.AccessSecret, 2)
	if err != nil {
		t.Fatal(err)
	}

	refreshToken, err := auth.CreateRefreshToken(data.Users[0], data.RefrshSecret, 2)
	if err != nil {
		t.Fatal(err)
	}

	IsAccessValide, err := auth.ValidateToken(accessToken, data.RefrshSecret)
	t.Log(IsAccessValide, err)
	IsRefreshValid, err := auth.ValidateToken(refreshToken, data.RefrshSecret)
	t.Log(IsRefreshValid, err)
}

func TestEncryption(t *testing.T) {
	gcm, keyString := encryption.CreateSymetricKey()
	t.Log(keyString)

	// encryption
	Message := "Hello there! what's going on the gc"
	// in the real case this message should be signed by the privet key
	encrypted := encryption.EncrypteSignedMessage(gcm, Message)
	t.Log(encrypted)

	// Decryption
	deccrypted := encryption.DecrypteMessage(gcm, encrypted)
	t.Log(encrypted, deccrypted)
}

func TestServer(t *testing.T) {
	data := testdata.GetData()
	server := socket.SocketServer{
		StorageFunc: data.StorageFunc,
	}

	engine := gin.Default()

	engine.GET("/Testws1", func(c *gin.Context) {
		room, err := server.GetRoom(data.Rooms[0])
		if err != nil {
			t.Fatal(err)
		}

		ws, err := socket.DefaultUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			t.Fatal(err)
		}

		client := socket.NewClient(&data.Users[1], ws, room)
		room.Join <- client

		go client.ReadIn()
		go client.WriteOut()
	})

	server.InitAndRun(data.Rooms)
	engine.Run()
}
