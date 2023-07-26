// Code generated by Kitex v0.6.2. DO NOT EDIT.

package pongservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	pong "simple-douyin/kitex_gen/pong"
)

func serviceInfo() *kitex.ServiceInfo {
	return pongServiceServiceInfo
}

var pongServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PongService"
	handlerType := (*pong.PongService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Pong": kitex.NewMethodInfo(pongHandler, newPongServicePongArgs, newPongServicePongResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "pong",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func pongHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*pong.PongServicePongArgs)
	realResult := result.(*pong.PongServicePongResult)
	success, err := handler.(pong.PongService).Pong(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPongServicePongArgs() interface{} {
	return pong.NewPongServicePongArgs()
}

func newPongServicePongResult() interface{} {
	return pong.NewPongServicePongResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Pong(ctx context.Context, req *pong.PingReq) (r *pong.PongResp, err error) {
	var _args pong.PongServicePongArgs
	_args.Req = req
	var _result pong.PongServicePongResult
	if err = p.c.Call(ctx, "Pong", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
