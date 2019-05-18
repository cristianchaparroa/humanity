package websocket

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	localAddr  = fakeAddr(1)
	remoteAddr = fakeAddr(2)
)

type fakeAddr int

func (a fakeAddr) Network() string {
	return "net"
}

func (a fakeAddr) String() string {
	return "str"
}

type fakeNetConn struct {
	io.Reader
	io.Writer
}

func (c fakeNetConn) Close() error                       { return nil }
func (c fakeNetConn) LocalAddr() net.Addr                { return localAddr }
func (c fakeNetConn) RemoteAddr() net.Addr               { return remoteAddr }
func (c fakeNetConn) SetDeadline(t time.Time) error      { return nil }
func (c fakeNetConn) SetReadDeadline(t time.Time) error  { return nil }
func (c fakeNetConn) SetWriteDeadline(t time.Time) error { return nil }

// TestResponseWriter ...
type TestResponseWriter struct {
	brw *bufio.ReadWriter
	http.ResponseWriter
}

// NewTestResponseWriter ...
func NewTestResponseWriter() *TestResponseWriter {
	r := bufio.NewReaderSize(strings.NewReader(""), 10)

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	brw := bufio.NewReadWriter(r, w)
	return &TestResponseWriter{brw: brw}
}

// Header ...
func (resp *TestResponseWriter) Header() http.Header {
	return http.Header{}
}

// WriteHeader ...
func (resp *TestResponseWriter) WriteHeader(statusCode int) {}

// Hijack ...
func (resp *TestResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	conn := fakeNetConn{strings.NewReader(""), &bytes.Buffer{}}
	return conn, resp.brw, nil
}

// WSFackeConn is a Mock for websocket connection.
type WSFackeConn struct{}

// WriteJSON mocks the write a json in the websock connection
func (m *WSFackeConn) WriteJSON(o interface{}) error { return nil }

// ReadJSON mocks read a json that cames from the connection
func (m *WSFackeConn) ReadJSON(o interface{}) error { return nil }

// Close mocks close a websocket connection
func (m *WSFackeConn) Close() error { return nil }
