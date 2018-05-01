package balance

//轮询算法（定义结构体，实现DoBalance方法，就实现了Balancer负载均衡接口）

import (
	"errors"
)

//程序启动，马上注册算法（init函数是程序启动最先执行的函数）
func init() {
	RegisterBalance("roundrobin",&RoundrobinBalance{})
}

type RoundrobinBalance struct {
	curIndex int
}

func (p *RoundrobinBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error) {
	lens := len(insts)
	//如果没传递参数，退出
	if lens <= 0 {
		err = errors.New("No balace instance")
		return
	}

	//如果索引大于实例个数。索引就为0（防止索引越界）
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	//实例数取余，+1取模可以循环下标，有轮询功能
	p.curIndex = (p.curIndex +1) % lens

	return
}

