package database

import (
	"strconv"
	"strings"
)

func LeaveRoom(gid uint32, pid uint32) error {
	sqlMutex.Lock()
	defer sqlMutex.Unlock()
	row := SQLite.QueryRow(
		`SELECT playerList FROM lm2_rooms WHERE gid = ?`,
		gid,
	)
	var playerList string
	err := row.Scan(&playerList)
	if err != nil {
		return err
	}
	players := strings.Split(playerList, ";")
	if playerList == "" {
		players = []string{}
	}
	newPlayers := make([]string, 0)
	for _, player := range players {
		if player != strconv.FormatUint(uint64(pid), 10) {
			newPlayers = append(newPlayers, player)
		}
	}
	_, err = SQLite.Exec(
		`UPDATE lm2_rooms SET playerList=?, participationCount=? WHERE gid=?`,
		strings.Join(newPlayers, ";"),
		len(newPlayers),
		gid,
	)
	if err != nil {
		return err
	}
	return nil
}
