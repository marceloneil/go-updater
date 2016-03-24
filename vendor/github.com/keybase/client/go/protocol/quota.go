// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/quota.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type VerifySessionRes struct {
	Uid       UID    `codec:"uid" json:"uid"`
	Sid       string `codec:"sid" json:"sid"`
	Generated int    `codec:"generated" json:"generated"`
	Lifetime  int    `codec:"lifetime" json:"lifetime"`
}

type VerifySessionArg struct {
	Session string `codec:"session" json:"session"`
}

type QuotaInterface interface {
	VerifySession(context.Context, string) (VerifySessionRes, error)
}

func QuotaProtocol(i QuotaInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.quota",
		Methods: map[string]rpc.ServeHandlerDescription{
			"verifySession": {
				MakeArg: func() interface{} {
					ret := make([]VerifySessionArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]VerifySessionArg)
					if !ok {
						err = rpc.NewTypeError((*[]VerifySessionArg)(nil), args)
						return
					}
					ret, err = i.VerifySession(ctx, (*typedArgs)[0].Session)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type QuotaClient struct {
	Cli rpc.GenericClient
}

func (c QuotaClient) VerifySession(ctx context.Context, session string) (res VerifySessionRes, err error) {
	__arg := VerifySessionArg{Session: session}
	err = c.Cli.Call(ctx, "keybase.1.quota.verifySession", []interface{}{__arg}, &res)
	return
}
