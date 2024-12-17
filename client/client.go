package client

import (
	"context"
	"net"
)

type Client struct {
    addr string
}

func NewClient(addr string) *Client {
    return &Client {
        addr: addr,
    }
}

func (c *Client) Set(ctx context.Context, key string, val string) error {
    conn, err := net.Dial("tcp", c.addr)
}
