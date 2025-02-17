// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: redirect/redirect.proto

package redirect

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationRedirectServiceRedirect = "/demo.redirect.RedirectService/Redirect"

type RedirectServiceHTTPServer interface {
	Redirect(context.Context, *RedirectRequest) (*RedirectReply, error)
}

func RegisterRedirectServiceHTTPServer(s *http.Server, srv RedirectServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/redirect", _RedirectService_Redirect0_HTTP_Handler(srv))
}

func _RedirectService_Redirect0_HTTP_Handler(srv RedirectServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RedirectRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationRedirectServiceRedirect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Redirect(ctx, req.(*RedirectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RedirectReply)
		return ctx.Result(200, reply)
	}
}

type RedirectServiceHTTPClient interface {
	Redirect(ctx context.Context, req *RedirectRequest, opts ...http.CallOption) (rsp *RedirectReply, err error)
}

type RedirectServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewRedirectServiceHTTPClient(client *http.Client) RedirectServiceHTTPClient {
	return &RedirectServiceHTTPClientImpl{client}
}

func (c *RedirectServiceHTTPClientImpl) Redirect(ctx context.Context, in *RedirectRequest, opts ...http.CallOption) (*RedirectReply, error) {
	var out RedirectReply
	pattern := "/api/v1/redirect"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationRedirectServiceRedirect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
