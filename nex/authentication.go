package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(0)
	globals.AuthenticationServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 1,
		Patch: 0,
	})
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("3861a9f8")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Luigi's Mansion: Dark Moon - Auth==")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID())
		fmt.Printf("Method ID: %d\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_LUIGISMANSION2_AUTHENTICATION_SERVER_PORT")))
}
