package bootstraps

import "github.com/youssefhmidi/E2E_encryptedConnection/database"

func InitDatabase(IsReleaseMode bool) database.SqliteDatabase {
	if IsReleaseMode {
		db := database.NewDB("./database/db/production.db")
		db.InitModels()
		return db
	}
	db := database.NewDB("./database/db/testingdb.db")
	db.InitModels()
	return db
}
