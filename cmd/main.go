package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/routes"
)

func main() {
	log.Println(`E2E_encrytedConnection is a backend for a chat app which is meant to store encrypted data, 
				the chat app client side is the responsible for encryption`)
	env := bootstraps.NewEnv(".env")
	DB := bootstraps.InitDatabase(env.IsReleaseMode)
	App := bootstraps.DefaultApp(DB)
	App.Env = env

	router := gin.Default()
	routes.SetupRoutes(router, App.Env, DB, App.SocketServer)

	router.Run()
}
