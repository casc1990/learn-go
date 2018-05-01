package balance

//随机算法（定义结构体，实现DoBalance方法，就实现了Balancer负载均衡接口）

import (
	"errors"
	"math/rand"
)

//程序启动，马上注册算法（init函数是程序启动最先执行的函数）
func init() {
	RegisterBalance("random",&RandomBalance{})
}

type RandomBalance struct {}

func (p *RandomBalance) DoBalance(insts []*Instance,key ...string) (inst *Instance,err error) {
	lens := len(insts)
	//如果没传递参数，退出
	if lens <= 0 {
		err = errors.New("No balace instance")
		return
	}
	//从传递的切片中随机选一个实例
	index := rand.Intn(lens-1)
	inst = insts[index]

	return
}