package nex

import (
	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
)

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
	/*matchmake_extension.NewCommonMatchmakeExtensionProtocol(globals.SecureServer)
	matchmaking.NewCommonMatchMakingProtocol(globals.SecureServer)
	matchmaking_ext.NewCommonMatchMakingExtProtocol(globals.SecureServer)
	nattraversal.NewCommonNATTraversalProtocol(globals.SecureServer)*/
}
