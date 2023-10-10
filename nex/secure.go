package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(0)
	globals.SecureServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 1,
		Patch: 0,
	})
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("3861a9f8")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Luigi's Mansion: Dark Moon - Secure==")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID())
		fmt.Printf("Method ID: %d\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_LUIGISMANSION2_SECURE_SERVER_PORT")))
}
