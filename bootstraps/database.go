package bootstraps

import "github.com/youssefhmidi/E2E_encryptedConnection/database"

func InitDatabase(location string) database.SqliteDatabase {
	db := database.NewDB(location)
	db.InitModels()
	return db
}
