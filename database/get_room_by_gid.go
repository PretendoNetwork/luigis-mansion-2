package database

import (
	"encoding/hex"
	"time"

	"github.com/PretendoNetwork/nex-go"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func GetRoomByGid(gid uint32) (*match_making_types.MatchmakeSession, error) {
	row := SQLite.QueryRow(
		`SELECT ownerPid, hostPid, flags, gameMode, matchmakeSystemType, participationCount, sessionKey, startedTime, playerList FROM lm2_rooms WHERE gid=?`,
		gid,
	)
	var ownerPid, hostPid, flags, gameMode, matchmakeSystemType, participationCount uint32
	var sessionKey, playerList string
	var startedTime time.Time

	err := row.Scan(&ownerPid, &hostPid, &flags, &gameMode, &matchmakeSystemType, &participationCount, &sessionKey, &startedTime, &playerList)
	if err != nil {
		return nil, err
	}
	m := &match_making_types.MatchmakeSession{}
	m.ID = gid
	m.OwnerPID = ownerPid
	m.HostPID = hostPid
	m.Flags = flags
	m.GameMode = gameMode
	m.MatchmakeSystemType = matchmakeSystemType
	m.ParticipationCount = participationCount
	k, _ := hex.DecodeString("00000000000000000000000000000000")
	m.SessionKey = k
	m.StartedTime = nex.NewDateTime(nex.NewDateTime(0).FromTimestamp(startedTime))

	return m, nil
}
