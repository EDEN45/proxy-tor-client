package srv

import (
	"github.com/EDEN45/proxy-tor-client/config"
	"h12.io/socks"
	"log"
	"net/http"
)

type Socks5ClientService struct {
	client *http.Client
	config config.Config
}

func NewSocks5ClientService(config config.Config) *Socks5ClientService {
	return &Socks5ClientService{
		config: config,
		client: initProxyClient(config),
	}
}

func initProxyClient(config config.Config) *http.Client {
	dialSocksProxy := socks.DialSocksProxy(socks.SOCKS5, config.UrlConnection)

	transport := &http.Transport{Dial: dialSocksProxy}
	return &http.Client{Transport: transport}
}

func (s *Socks5ClientService) DoRequest(r *http.Request) (response *http.Response, err error) {
	return s.client.Do(r)
}

func (s *Socks5ClientService) HttpGet(url string) (response *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("User-Agent", s.config.UserAgents[0])
	return s.client.Do(req)
}
