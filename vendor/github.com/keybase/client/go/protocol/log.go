// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/log.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type RegisterLoggerArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	Name      string   `codec:"name" json:"name"`
	Level     LogLevel `codec:"level" json:"level"`
}

type LogInterface interface {
	RegisterLogger(context.Context, RegisterLoggerArg) error
}

func LogProtocol(i LogInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.log",
		Methods: map[string]rpc.ServeHandlerDescription{
			"registerLogger": {
				MakeArg: func() interface{} {
					ret := make([]RegisterLoggerArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RegisterLoggerArg)
					if !ok {
						err = rpc.NewTypeError((*[]RegisterLoggerArg)(nil), args)
						return
					}
					err = i.RegisterLogger(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LogClient struct {
	Cli rpc.GenericClient
}

func (c LogClient) RegisterLogger(ctx context.Context, __arg RegisterLoggerArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.log.registerLogger", []interface{}{__arg}, nil)
	return
}
