package database

import "github.com/PretendoNetwork/luigis-mansion-2/globals"

func InitSQLite() {
	var err error

	_, err = SQLite.Exec(`CREATE TABLE IF NOT EXISTS lm2_rooms (
		gid integer PRIMARY KEY,
		ownerPid integer DEFAULT 0,
		hostPid integer DEFAULT 0,
		flags integer DEFAULT 0,
		gameMode integer DEFAULT 0,
		matchmakeSystemType integer DEFAULT 0,
		participationCount integer DEFAULT 0,
		sessionKey text DEFAULT "",
		startedTime datetime DEFAULT "",
		playerList string DEFAULT ""
	);`)
	if err != nil {
		globals.Logger.Critical(err.Error())
		return
	}

	globals.Logger.Success("SQLite tables created")
}
