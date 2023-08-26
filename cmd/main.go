package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"
	"github.com/youssefhmidi/E2E_encryptedConnection/routes"
)

func main() {
	DB := bootstraps.InitDatabase("./database/db/testingdb.db")
	App := bootstraps.DefaultApp("../.env", DB)
	router := gin.Default()

	routes.SetupRoutes(router, App.Env, DB, App.SocketServer)

	router.Run()
}
