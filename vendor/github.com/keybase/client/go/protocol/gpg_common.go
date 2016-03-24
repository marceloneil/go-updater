// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/gpg_common.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type GPGKey struct {
	Algorithm  string        `codec:"algorithm" json:"algorithm"`
	KeyID      string        `codec:"keyID" json:"keyID"`
	Creation   string        `codec:"creation" json:"creation"`
	Expiration string        `codec:"expiration" json:"expiration"`
	Identities []PGPIdentity `codec:"identities" json:"identities"`
}

type GpgCommonInterface interface {
}

func GpgCommonProtocol(i GpgCommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.gpgCommon",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GpgCommonClient struct {
	Cli rpc.GenericClient
}
