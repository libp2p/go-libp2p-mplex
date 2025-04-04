package mplex

import (
	"context"

	"github.com/libp2p/go-libp2p/core/network"

	mp "github.com/libp2p/go-mplex"
)

type conn mp.Multiplex

var _ network.MuxedConn = &conn{}

// NewMuxedConn constructs a new Conn from a *mp.Multiplex.
func NewMuxedConn(m *mp.Multiplex) network.MuxedConn {
	return (*conn)(m)
}

func (c *conn) Close() error {
	return c.mplex().Close()
}

// CloseWithError cloes the connection ignoring the error code.
// There's no mechanism to send an error code to the peer in mplex.
func (c *conn) CloseWithError(_ network.ConnErrorCode) error {
	return c.mplex().Close()
}

func (c *conn) IsClosed() bool {
	return c.mplex().IsClosed()
}

// OpenStream creates a new stream.
func (c *conn) OpenStream(ctx context.Context) (network.MuxedStream, error) {
	s, err := c.mplex().NewStream(ctx)
	if err != nil {
		return nil, err
	}
	return (*stream)(s), nil
}

// AcceptStream accepts a stream opened by the other side.
func (c *conn) AcceptStream() (network.MuxedStream, error) {
	s, err := c.mplex().Accept()
	if err != nil {
		return nil, err
	}
	return (*stream)(s), nil
}

func (c *conn) mplex() *mp.Multiplex {
	return (*mp.Multiplex)(c)
}
