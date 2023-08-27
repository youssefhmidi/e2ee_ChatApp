package routes

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/_internals/socket"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/controllers"
	"github.com/youssefhmidi/E2E_encryptedConnection/database"
	"github.com/youssefhmidi/E2E_encryptedConnection/middlewares"
	"github.com/youssefhmidi/E2E_encryptedConnection/repository"
	"github.com/youssefhmidi/E2E_encryptedConnection/services"
)

func SetupRoutes(engine *gin.Engine, env *bootstraps.Env, db database.SqliteDatabase, socketServer *socket.SocketServer) {
	// setting up the repositories
	ur := repository.NewUserRepository(db)
	mr := repository.NewMessageRepository(db)
	cr := repository.NewChatRepository(db)

	// setting up internal services
	Secrets := map[string]string{
		"access":  env.AccessTokenSecret,
		"refresh": env.RefreshTokenSecret,
	}
	Expiry := map[string]int{
		"access":  env.AccessTokenExpiry,
		"refresh": env.RefreshTokenExpiry,
	}
	jwtS := services.NewJwtService(Secrets, Expiry)

	// initializing the services
	us := services.NewUserService(ur, jwtS)
	ls := services.NewLoginService(ur, jwtS)
	ss := services.NewSignupService(ur, jwtS)
	rs := services.NewRoomService(ur, cr)
	cs := services.NewChatService(mr)
	wss := services.NewWebsocketService(cs, rs)

	// initializing the controllers
	uc := controllers.NewUserController(env, us)
	lc := controllers.NewLoginController(env, ls)
	sc := controllers.NewSignupController(env, ss)
	rc := controllers.NewRoomController(socketServer, wss, rs, us, env)

	// setting up the server Storagefunc and runnig it
	socketServer.StorageFunc = wss.StoreMsgsToDatabase(context.Background())
	rooms, err := cr.GetRooms()
	if err != nil {
		log.Fatal(err)
	}
	socketServer.InitAndRun(rooms)

	// adding endpoints

	// '/login', '/signup', '/refresh'
	newLoginRoute(engine, lc)
	newSignupRoute(engine, sc)
	newRefreshRoute(engine, uc, Secrets["refresh"])

	// '/users/@me'
	userGroup := engine.Group("/users")
	userGroup.Use(middlewares.UseTokenVerification(Secrets["access"], "access"))
	newUserRoute(userGroup, uc)

	// '/chat/'
	roomGroup := engine.Group("/chat")
	roomGroup.Use(middlewares.UseWebsocketAuth(Secrets["access"]))
	newRoomRoutes(roomGroup, rc)
}
