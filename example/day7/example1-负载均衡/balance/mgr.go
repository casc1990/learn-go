package balance

import "fmt"

//负载均衡算法管理类,提供扩展功能

type BalanceMgr struct {
	//存储所有的负载均衡算法。amp的key是算法名称，value是实现了Balance方法的接口
	allBalance map[string]Balancer
}
//map对象需要使用make实例化
var mgr = BalanceMgr{
	allBalance: make(map[string]Balancer),
}

//算法注册(将算法存到map)
func (p *BalanceMgr) registerBalance(name string,b Balancer) {
	p.allBalance[name] = b
}

//向外只用暴露注册方法(RegisterBalance)
func RegisterBalance(name string,b Balancer) {
	mgr.registerBalance(name,b)  //调用registerBalance方法将用户的方法注册
}

//向外提供负载均衡方法(封装了底层结构体的DoBalance方法)
func DoBalance(name string,insts []*Instance) (inst *Instance,err error) {
	//判断allBalance里是否有用户指定的方法
	balancer,ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer!",name)
		return
	}
	fmt.Printf("use %s balancer\n",name)
	inst,err = balancer.DoBalance(insts)
	return
}
