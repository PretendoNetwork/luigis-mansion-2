package matchmake_extension

import (
	"encoding/hex"
	"fmt"

	"github.com/PretendoNetwork/luigis-mansion-2/database"
	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	"github.com/PretendoNetwork/nex-go"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
)

var MODES []string = []string{"Hunter", "Rush", "Polterpup"}
var DIFFS []string = []string{"Normal", "Hard", "Expert"}
var FLOORS map[int]string = map[int]string{
	5:   "5 Floors",
	10:  "10 Floors",
	25:  "25 Floors",
	100: "Endless",
}

func CreateMatchmakeSession(err error, client *nex.Client, callID uint32, data *nex.DataHolder, strMessage string, participationCount uint16) uint32 {
	matchmakeSession := data.ObjectData().(*match_making_types.MatchmakeSession)
	fmt.Printf("%b\n", matchmakeSession.GameMode)

	gid, _ := database.CreateRoom(client.PID(), matchmakeSession.Flags, matchmakeSession.GameMode, matchmakeSession.MatchmakeSystemType)
	database.JoinRoom(gid, client.PID())

	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteUInt32LE(gid)
	b, _ := hex.DecodeString("00000000000000000000000000000000")
	rmcResponseStream.WriteBuffer(b)

	rmcResponseBody := rmcResponseStream.Bytes()

	// Build response packet
	rmcResponse := nex.NewRMCResponse(matchmake_extension.ProtocolID, callID)
	rmcResponse.SetSuccess(matchmake_extension.MethodCreateMatchmakeSession, rmcResponseBody)

	rmcResponseBytes := rmcResponse.Bytes()

	responsePacket, _ := nex.NewPacketV1(client, nil)

	responsePacket.SetVersion(1)
	responsePacket.SetSource(0xA1)
	responsePacket.SetDestination(0xAF)
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	globals.SecureServer.Send(responsePacket)

	return 0
}
