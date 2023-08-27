package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/routes"
)

func main() {
	log.Println("E2E_encrytedConnection is a backend for a chat app which is meant to store encrypted data, the chat app client side is the responsible for encryption")
	DB := bootstraps.InitDatabase("./database/db/testingdb.db")
	App := bootstraps.DefaultApp(".env", DB)
	router := gin.Default()

	routes.SetupRoutes(router, App.Env, DB, App.SocketServer)

	router.Run()
}
