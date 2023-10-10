package nex

import (
	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	lm2_match_making_ext "github.com/PretendoNetwork/luigis-mansion-2/nex/match_making_ext"
	lm2_matchmake_extension "github.com/PretendoNetwork/luigis-mansion-2/nex/matchmake_extension"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/match-making-ext"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
)

func registerSecureServerNEXProtocols() {
	matchmake_extension_proto := matchmake_extension.NewProtocol(globals.SecureServer)
	matchmake_extension_proto.BrowseMatchmakeSession(lm2_matchmake_extension.BrowseMatchmakeSession)
	matchmake_extension_proto.CreateMatchmakeSession(lm2_matchmake_extension.CreateMatchmakeSession)
	matchmake_extension_proto.OpenParticipation(lm2_matchmake_extension.OpenParticipation)
	match_making_ext_proto := match_making_ext.NewProtocol(globals.SecureServer)
	match_making_ext_proto.EndParticipation(lm2_match_making_ext.EndParticipation)
}
