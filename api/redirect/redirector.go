package redirect

import "github.com/go-kratos/kratos/v2/transport/http"

var _ http.Redirector = (*RedirectReply)(nil)

func (s *RedirectReply) Redirect() (string, int) {
	return s.RedirectTo, int(s.StatusCode)
}
