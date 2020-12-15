package ss

import (
	"net"
	"time"

	"github.com/chinaboard/coral/core/proxy"

	"github.com/chinaboard/coral/config"

	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)

type ShadowsocksProxy struct {
	name    string
	Timeout time.Duration
	Cipher  *ss.Cipher
	Address string
}

func New(server config.CoralServer) (proxy.Proxy, error) {
	cipher, err := ss.NewCipher(server.Method, server.Password)
	if err != nil {
		return nil, err
	}

	return &ShadowsocksProxy{
		name:    server.Name,
		Timeout: server.ReadTimeout,
		Cipher:  cipher,
		Address: server.Address(),
	}, nil
}

func (this *ShadowsocksProxy) Dial(network, addr string) (net.Conn, time.Duration, error) {
	conn, err := ss.Dial(addr, this.Address, this.Cipher.Copy())
	return conn, this.Timeout, err
}

func (this *ShadowsocksProxy) Name() string {
	return this.name
}

func (this *ShadowsocksProxy) Direct() bool {
	return false
}
