package database

import (
	"encoding/hex"
	"time"

	"github.com/PretendoNetwork/nex-go"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func GetRooms() ([]*match_making_types.MatchmakeSession, error) {
	m := make([]*match_making_types.MatchmakeSession, 0)
	rows, _ := SQLite.Query(
		`SELECT gid, ownerPid, hostPid, flags, gameMode, matchmakeSystemType, participationCount, sessionKey, startedTime, playerList FROM lm2_rooms`,
	)
	var gid, ownerPid, hostPid, flags, gameMode, matchmakeSystemType, participationCount uint32
	var sessionKey, playerList string
	var startedTime time.Time

	for rows.Next() {
		err := rows.Scan(&gid, &ownerPid, &hostPid, &flags, &gameMode, &matchmakeSystemType, &participationCount, &sessionKey, &startedTime, &playerList)
		if err != nil {
			return nil, err
		}
		session := match_making_types.NewMatchmakeSession()
		session.ID = gid
		session.OwnerPID = ownerPid
		session.HostPID = hostPid
		session.Flags = flags
		session.GameMode = gameMode
		session.MatchmakeSystemType = matchmakeSystemType
		session.ParticipationCount = participationCount
		k, _ := hex.DecodeString("00000000000000000000000000000000")
		session.SessionKey = k
		session.StartedTime = nex.NewDateTime(nex.NewDateTime(0).FromTimestamp(startedTime))

		m = append(m, session)
	}

	return m, nil
}
