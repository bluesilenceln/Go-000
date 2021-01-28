package echo

import (
	"bufio"
	"context"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
	addr string

	rd *bufio.Reader
	wd *bufio.Writer

	ch chan string
}

func NewClient(addr string) *Client {
	c := new(Client)
	c.addr = addr
	c.ch = make(chan string)

	return c
}

func (c *Client) Start(ctx context.Context) error {
	var err error
	c.conn, err = net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	c.rd = bufio.NewReader(c.conn)
	c.wd = bufio.NewWriter(c.conn)
	go c.handleRead()

	return nil
}

func (c *Client) Stop(ctx context.Context) error {
	close(c.ch)
	return c.conn.Close()
}

func (c *Client) handleRead() {
	for {
		msg, err := c.rd.ReadString('\n')
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("read error: %v", err)
			}
			break
		}
		c.ch <- msg
	}
}

func (c *Client) Send(msg string) error {
	_, err := c.wd.WriteString(msg)
	if err != nil {
		return err
	}

	err = c.wd.Flush()
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Recv() <- chan string {
	return c.ch
}

