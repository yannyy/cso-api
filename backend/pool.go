package backend

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrorClosed  = errors.New("connection closed")
	ErrorTimeout = errors.New("connection pool timeout")
)

var timers = sync.Pool{
	New: func() interface{} {
		t := time.NewTimer(time.Hour)
		t.Stop()
		return t
	},
}

type Conn struct {
}

type ConnPoolOption struct {
	OnClose func(*Conn) error

	PoolSize           int
	MinIdleConns       int
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
	PoolTimeout        time.Duration
	MaxConnAge         time.Duration
}

type AddConnFunc func() interface{}

type ConnPool struct {
	opt *ConnPoolOption

	connsMu sync.Mutex

	conns        []interface{}
	idleConns    []interface{}
	poolSize     int
	idelConnSize int
	addConnFunc  AddConnFunc
}

func (c *ConnPool) Lock() {
	c.connsMu.Lock()
}

func (c *ConnPool) Unlock() {
	c.connsMu.Unlock()
}

func NewConnPool(opt *ConnPoolOption, addConnFunc AddConnFunc) *ConnPool {
	p := &ConnPool{
		opt:         opt,
		conns:       make([]interface{}, 0),
		idleConns:   make([]interface{}, 0),
		addConnFunc: addConnFunc,
	}

	p.checkMinIdelConns()

	return p
}

func (p *ConnPool) checkMinIdelConns() {
	p.Lock()
	if p.opt.MinIdleConns == 0 {
		return
	}
	for p.poolSize < p.opt.PoolSize && p.idelConnSize < p.opt.MinIdleConns {
		p.poolSize++
		p.idelConnSize++
		conn := p.addConnFunc()
		p.conns = append(p.conns, conn)
		p.idleConns = append(p.idleConns, conn)
	}
	p.Unlock()
}

func (p *ConnPool) Get() interface{} {
	return p.conns[0]
}
