// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/track.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type TrackArg struct {
	SessionID        int          `codec:"sessionID" json:"sessionID"`
	UserAssertion    string       `codec:"userAssertion" json:"userAssertion"`
	Options          TrackOptions `codec:"options" json:"options"`
	ForceRemoteCheck bool         `codec:"forceRemoteCheck" json:"forceRemoteCheck"`
}

type TrackWithTokenArg struct {
	SessionID  int          `codec:"sessionID" json:"sessionID"`
	TrackToken TrackToken   `codec:"trackToken" json:"trackToken"`
	Options    TrackOptions `codec:"options" json:"options"`
}

type UntrackArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

type CheckTrackingArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type FakeTrackingChangedArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Username  string `codec:"username" json:"username"`
}

type TrackInterface interface {
	Track(context.Context, TrackArg) error
	TrackWithToken(context.Context, TrackWithTokenArg) error
	Untrack(context.Context, UntrackArg) error
	CheckTracking(context.Context, int) error
	FakeTrackingChanged(context.Context, FakeTrackingChangedArg) error
}

func TrackProtocol(i TrackInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.track",
		Methods: map[string]rpc.ServeHandlerDescription{
			"track": {
				MakeArg: func() interface{} {
					ret := make([]TrackArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TrackArg)
					if !ok {
						err = rpc.NewTypeError((*[]TrackArg)(nil), args)
						return
					}
					err = i.Track(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"trackWithToken": {
				MakeArg: func() interface{} {
					ret := make([]TrackWithTokenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TrackWithTokenArg)
					if !ok {
						err = rpc.NewTypeError((*[]TrackWithTokenArg)(nil), args)
						return
					}
					err = i.TrackWithToken(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"untrack": {
				MakeArg: func() interface{} {
					ret := make([]UntrackArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]UntrackArg)
					if !ok {
						err = rpc.NewTypeError((*[]UntrackArg)(nil), args)
						return
					}
					err = i.Untrack(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"checkTracking": {
				MakeArg: func() interface{} {
					ret := make([]CheckTrackingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CheckTrackingArg)
					if !ok {
						err = rpc.NewTypeError((*[]CheckTrackingArg)(nil), args)
						return
					}
					err = i.CheckTracking(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"fakeTrackingChanged": {
				MakeArg: func() interface{} {
					ret := make([]FakeTrackingChangedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FakeTrackingChangedArg)
					if !ok {
						err = rpc.NewTypeError((*[]FakeTrackingChangedArg)(nil), args)
						return
					}
					err = i.FakeTrackingChanged(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TrackClient struct {
	Cli rpc.GenericClient
}

func (c TrackClient) Track(ctx context.Context, __arg TrackArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.track.track", []interface{}{__arg}, nil)
	return
}

func (c TrackClient) TrackWithToken(ctx context.Context, __arg TrackWithTokenArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.track.trackWithToken", []interface{}{__arg}, nil)
	return
}

func (c TrackClient) Untrack(ctx context.Context, __arg UntrackArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.track.untrack", []interface{}{__arg}, nil)
	return
}

func (c TrackClient) CheckTracking(ctx context.Context, sessionID int) (err error) {
	__arg := CheckTrackingArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.track.checkTracking", []interface{}{__arg}, nil)
	return
}

func (c TrackClient) FakeTrackingChanged(ctx context.Context, __arg FakeTrackingChangedArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.track.fakeTrackingChanged", []interface{}{__arg}, nil)
	return
}
