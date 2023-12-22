package main

import (
	"math/rand"
	"sync"
)

const (
	ON  = 1
	OFF = 0
)

type Conn struct {
	ID     int64
	Target string
	Status int
}

func NewCOnn(target string) *Conn {
	return &Conn{
		ID:     rand.Int63(),
		Target: target,
		Status: ON,
	}
}

func (c Conn) GetStatus() int {
	return c.Status
}

type ConnPool struct {
	sync.Pool
}

func GetPool(target string) (*ConnPool, error) {
	return &ConnPool{Pool: sync.Pool{
		New: func() any {
			return NewCOnn(target)
		},
	}}, nil
}

func (c *ConnPool) Get() *Conn {
	conn := c.Pool.Get().(*Conn)
	if conn.GetStatus() == OFF {
		conn = c.Pool.New().(*Conn)
	}
	return conn
}

func main() {

}