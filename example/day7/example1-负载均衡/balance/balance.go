package balance

//定义负载均衡接口（Balancer），包含DoBalance方法(实现了这个方法，就实现了这个接口)

type Balancer interface {
	//DoBalance方法接受切片类型的实例列表和可变长的string类型（string类型用于我们实现hash算法）
	DoBalance([]*Instance, ...string) (*Instance,error)
}
