package srv

import "net/http"

type Socks5ClientContract interface {
	DoRequest(r *http.Request) (response *http.Response, err error)
	HttpGet(url string) (response *http.Response, err error)
}
