package mplex

import (
	"time"

	"github.com/libp2p/go-libp2p/core/network"

	mp "github.com/libp2p/go-mplex"
)

// stream implements network.MuxedStream over mplex.Stream.
type stream mp.Stream

var _ network.MuxedStream = &stream{}

func (s *stream) Read(b []byte) (n int, err error) {
	n, err = s.mplex().Read(b)
	if err == mp.ErrStreamReset {
		err = network.ErrReset
	}

	return n, err
}

func (s *stream) Write(b []byte) (n int, err error) {
	n, err = s.mplex().Write(b)
	if err == mp.ErrStreamReset {
		err = network.ErrReset
	}

	return n, err
}

func (s *stream) Close() error {
	return s.mplex().Close()
}

func (s *stream) CloseWrite() error {
	return s.mplex().CloseWrite()
}

func (s *stream) CloseRead() error {
	return s.mplex().CloseRead()
}

func (s *stream) Reset() error {
	return s.mplex().Reset()
}

// ResetWithError resets the stream ignoring the error code.
// There's no mechanism to send an error code to the peer in mplex.
func (s *stream) ResetWithError(_ network.StreamErrorCode) error {
	return s.mplex().Reset()
}

func (s *stream) SetDeadline(t time.Time) error {
	return s.mplex().SetDeadline(t)
}

func (s *stream) SetReadDeadline(t time.Time) error {
	return s.mplex().SetReadDeadline(t)
}

func (s *stream) SetWriteDeadline(t time.Time) error {
	return s.mplex().SetWriteDeadline(t)
}

func (s *stream) mplex() *mp.Stream {
	return (*mp.Stream)(s)
}
