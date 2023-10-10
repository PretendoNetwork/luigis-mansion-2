package database

import (
	"errors"
	"strconv"
	"strings"

	"github.com/PretendoNetwork/luigis-mansion-2/globals"
)

func JoinRoom(gid uint32, pid uint32) error {
	sqlMutex.Lock()
	defer sqlMutex.Unlock()
	row := SQLite.QueryRow(
		`SELECT participationCount, playerList FROM lm2_rooms WHERE gid = ?`,
		gid,
	)
	var participationCount uint32
	var playerList string
	err := row.Scan(&participationCount, &playerList)
	if err != nil {
		return err
	}
	if participationCount >= globals.MAX_PLAYERS {
		return errors.New("cannot join room, room is full")
	}
	players := strings.Split(playerList, ";")
	if playerList == "" {
		players = []string{}
	}
	players = append(players, strconv.FormatUint(uint64(pid), 10))
	_, err = SQLite.Exec(
		`UPDATE lm2_rooms SET playerList=?, participationCount=? WHERE gid=?`,
		strings.Join(players, ";"),
		len(players),
		gid,
	)
	if err != nil {
		return err
	}
	return nil
}
