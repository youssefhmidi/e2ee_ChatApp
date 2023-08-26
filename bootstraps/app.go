package bootstraps

import (
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/database"
)

// the app structure represent an app wich contain the db , the env  and the socket server
type App struct {
	// an interface for accessing a database
	Db database.SqliteDatabase
	// for getting all the envirement variables
	Env *Env
	// for runing the socket server for the chat service
	SocketServer *socket.SocketServer
}

// an initializer for making a basic App
// env : location of the '.env' file
// db : usualy returned from the InitDatabase func
func DefaultApp(env string, db database.SqliteDatabase) *App {
	return &App{
		Db:           db,
		Env:          NewEnv(env),
		SocketServer: &socket.SocketServer{},
	}
}
