// Code generated by Kitex v0.9.1. DO NOT EDIT.

package miscservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	sc_misc_api "github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ReaderLogin": kitex.NewMethodInfo(
		readerLoginHandler,
		newMiscServiceReaderLoginArgs,
		newMiscServiceReaderLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"EditorLogin": kitex.NewMethodInfo(
		editorLoginHandler,
		newMiscServiceEditorLoginArgs,
		newMiscServiceEditorLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	miscServiceServiceInfo                = NewServiceInfo()
	miscServiceServiceInfoForClient       = NewServiceInfoForClient()
	miscServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return miscServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return miscServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return miscServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "MiscService"
	handlerType := (*sc_misc_api.MiscService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "sc_misc_api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func readerLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sc_misc_api.MiscServiceReaderLoginArgs)
	realResult := result.(*sc_misc_api.MiscServiceReaderLoginResult)
	success, err := handler.(sc_misc_api.MiscService).ReaderLogin(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMiscServiceReaderLoginArgs() interface{} {
	return sc_misc_api.NewMiscServiceReaderLoginArgs()
}

func newMiscServiceReaderLoginResult() interface{} {
	return sc_misc_api.NewMiscServiceReaderLoginResult()
}

func editorLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sc_misc_api.MiscServiceEditorLoginArgs)
	realResult := result.(*sc_misc_api.MiscServiceEditorLoginResult)
	success, err := handler.(sc_misc_api.MiscService).EditorLogin(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMiscServiceEditorLoginArgs() interface{} {
	return sc_misc_api.NewMiscServiceEditorLoginArgs()
}

func newMiscServiceEditorLoginResult() interface{} {
	return sc_misc_api.NewMiscServiceEditorLoginResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ReaderLogin(ctx context.Context, request *sc_misc_api.ReaderLoginRequest) (r *sc_misc_api.ReaderLoginResponse, err error) {
	var _args sc_misc_api.MiscServiceReaderLoginArgs
	_args.Request = request
	var _result sc_misc_api.MiscServiceReaderLoginResult
	if err = p.c.Call(ctx, "ReaderLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EditorLogin(ctx context.Context, request *sc_misc_api.EditorLoginRequest) (r *sc_misc_api.EditorLoginResponse, err error) {
	var _args sc_misc_api.MiscServiceEditorLoginArgs
	_args.Request = request
	var _result sc_misc_api.MiscServiceEditorLoginResult
	if err = p.c.Call(ctx, "EditorLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}