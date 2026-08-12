package http

import (
	"context"
	"net"

	"github.com/lzpls/tls"
)

type TLSConn interface {
	net.Conn
	NetConn() net.Conn
	HandshakeContext(ctx context.Context) error
	ConnectionState() tls.ConnectionState
}

var _ TLSConn = (*tls.Conn)(nil)
