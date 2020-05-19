package gcd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/url"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const writeSize = 15000000 // 15mb

// adapted from https://github.com/chromedp/chromedp/blob/8e0a16689423d48d8907c62a543c7ea468059228/conn.go
type wsConn struct {
	conn   net.Conn
	writer wsutil.Writer
}

func newWsConnDial(ctx context.Context, url string) (*wsConn, error) {
	wconn := &wsConn{}
	conn, br, _, err := ws.Dial(ctx, url)
	if err != nil {
		return nil, err
	}
	if br != nil {
		panic("br should be nil")
	}
	wconn.conn = conn
	wconn.writer = *wsutil.NewWriterBufferSize(conn, ws.StateClientSide, ws.OpText, 1<<20)
	return wconn, nil
}

func formatURL(toFormat string) string {
	u, err := url.Parse(toFormat)
	if err != nil {
		return ""
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return ""
	}
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return ""
	}
	u.Host = net.JoinHostPort(addr.IP.String(), port)
	return u.String()
}

func (c *wsConn) Read(ctx context.Context, msg *[]byte) error {
	// get websocket reader
	c.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	reader := wsutil.NewReader(c.conn, ws.StateClientSide)
	h, err := reader.NextFrame()
	if err != nil {
		return err
	}

	if h.OpCode == ws.OpClose {
		return io.EOF
	}

	if h.OpCode != ws.OpText {
		return fmt.Errorf("InvalidWebsocketMessage")
	}

	var b bytes.Buffer
	if _, err := b.ReadFrom(reader); err != nil {
		return err
	}

	*msg = b.Bytes()
	return nil
}

// Write writes a message.
func (c *wsConn) Write(_ context.Context, msg []byte) error {
	c.writer.Reset(c.conn, ws.StateClientSide, ws.OpText)
	if _, err := c.writer.Write(msg); err != nil {
		return err
	}
	return c.writer.Flush()
}

func (c *wsConn) Close() error {
	return c.Close()
}
