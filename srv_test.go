package proxy_tor_client

import (
	"github.com/EDEN45/proxy-tor-client/srv"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	service := srv.NewService()

	resp, err := service.HttpGet("https://google.com")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("Status not OK")
	}
}
