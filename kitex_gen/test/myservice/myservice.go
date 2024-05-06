// Code generated by Kitex v0.9.1. DO NOT EDIT.

package myservice

import (
	test "code.howimetmrright.com/smart-classroom/sc-misc-idl/kitex_gen/test"
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Hello": kitex.NewMethodInfo(
		helloHandler,
		newMyServiceHelloArgs,
		newMyServiceHelloResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	myServiceServiceInfo                = NewServiceInfo()
	myServiceServiceInfoForClient       = NewServiceInfoForClient()
	myServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return myServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return myServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return myServiceServiceInfoForClient
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
	serviceName := "MyService"
	handlerType := (*test.MyService)(nil)
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
		"PackageName": "test",
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

func helloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*test.MyServiceHelloArgs)
	realResult := result.(*test.MyServiceHelloResult)
	success, err := handler.(test.MyService).Hello(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newMyServiceHelloArgs() interface{} {
	return test.NewMyServiceHelloArgs()
}

func newMyServiceHelloResult() interface{} {
	return test.NewMyServiceHelloResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Hello(ctx context.Context, req *test.MyReq) (r string, err error) {
	var _args test.MyServiceHelloArgs
	_args.Req = req
	var _result test.MyServiceHelloResult
	if err = p.c.Call(ctx, "Hello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
