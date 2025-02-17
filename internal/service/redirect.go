package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-demo/api/redirect"
)

var _ redirect.RedirectServiceServer = (*RedirectService)(nil)

type RedirectService struct {
	log *log.Helper
	redirect.UnimplementedRedirectServiceServer
}

func (r RedirectService) Redirect(ctx context.Context, request *redirect.RedirectRequest) (*redirect.RedirectReply, error) {
	return &redirect.RedirectReply{
		RedirectTo: "https://www.baidu.com",
		StatusCode: 302,
	}, nil

}

func NewRedirectService(logger log.Logger) *RedirectService {
	return &RedirectService{log: log.NewHelper(logger)}
}
