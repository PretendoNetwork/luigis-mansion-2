package database

import "time"

func CreateRoom(ownerPid uint32, flags uint32, gameMode uint32, matchmakeSystemType uint32) (uint32, error) {
	sqlMutex.Lock()
	defer sqlMutex.Unlock()
	_, err := SQLite.Exec(
		`INSERT INTO lm2_rooms (
			ownerPid, hostPid, flags, gameMode, matchmakeSystemType, startedTime
		) VALUES (
			?, ?, ?, ?, ?, ?
		)`,
		ownerPid, ownerPid, flags, gameMode, matchmakeSystemType, time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return 0, err
	}
	row := SQLite.QueryRow(
		`SELECT last_insert_rowid()`,
	)
	var gid uint32
	err = row.Scan(&gid)
	if err != nil {
		return 0, err
	}

	return gid, nil
}
