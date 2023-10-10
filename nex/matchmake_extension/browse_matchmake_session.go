package matchmake_extension

import (
	"fmt"

	"github.com/PretendoNetwork/luigis-mansion-2/database"
	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	"github.com/PretendoNetwork/nex-go"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
)

func BrowseMatchmakeSession(err error, client *nex.Client, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32 {
	fmt.Println(searchCriteria)
	fmt.Println(resultRange)

	rooms, _ := database.GetRooms()
	dataRooms := make([]*nex.DataHolder, 0)

	for _, r := range rooms {
		d := nex.NewDataHolder()
		d.SetObjectData(r)
		dataRooms = append(dataRooms, d)
	}

	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)

	rmcResponseStream.WriteListDataHolder(dataRooms)

	rmcResponseBody := rmcResponseStream.Bytes()

	fmt.Printf("%x\n", rmcResponseBody)

	// Build response packet
	rmcResponse := nex.NewRMCResponse(matchmake_extension.ProtocolID, callID)
	rmcResponse.SetSuccess(matchmake_extension.MethodBrowseMatchmakeSession, rmcResponseBody)

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
