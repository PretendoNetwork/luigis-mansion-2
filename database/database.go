package database

import (
	"database/sql"
	"sync"

	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	_ "github.com/mattn/go-sqlite3"
)

var SQLite *sql.DB
var sqlMutex *sync.Mutex = &sync.Mutex{}

func Connect() {
	var err error
	SQLite, err = sql.Open("sqlite3", "lm2.db")

	if err != nil {
		globals.Logger.Critical(err.Error())
	}

	InitSQLite()
	globals.Logger.Success("Initialized SQLite!")
}
