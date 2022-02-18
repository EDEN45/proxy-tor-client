package srv

import "github.com/EDEN45/proxy-tor-client/config"

func NewService() Socks5ClientContract {
	conf := config.Boot()
	return NewSocks5ClientService(conf)
}
