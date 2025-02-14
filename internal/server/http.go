package server

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"kratos-demo/api/file_upload"
	"kratos-demo/internal/conf"
	"kratos-demo/internal/service"
	"kratos-demo/pkg/my_encode"
	"kratos-demo/pkg/server/op/http_encoder"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.FileUploadService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.ResponseEncoder(http_encoder.MyResponseEncoder),
		http.ErrorEncoder(http_encoder.MyErrorResponseEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	encoding.RegisterCodec(my_encode.FormCodec{})
	srv := http.NewServer(opts...)
	file_upload.RegisterFileUploadServiceHTTPServer(srv, greeter)
	return srv
}
