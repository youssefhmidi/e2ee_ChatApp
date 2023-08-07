package main

import "github.com/youssefhmidi/E2E_encryptedConnection/bootstraps"

func main() {
	bootstraps.InitDatabase("./database/db/testingdb.db")
}
