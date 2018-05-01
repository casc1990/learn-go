package balance

//定义负载均衡后端实例

import (
	"strconv"
)

type Instance struct {
	host string
	port int
}

func (p *Instance) GetHost() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}

func NewInstance(host string,port int) *Instance {
	return &Instance{
		host:host,
		port:port,
	}
}