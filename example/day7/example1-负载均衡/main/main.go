package main

import (
	"learn-go/example/day7/example1-负载均衡/balance"
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	//定义实例
	insts := []*balance.Instance{}
	for i :=0;i<10;i++ {
		host := fmt.Sprintf("192.168.%d.%d",rand.Intn(254),rand.Intn(254))
		inst := balance.NewInstance(host,8080)
		insts = append(insts,inst)
	}


	//调用用户指定的算法
	var balancer balance.Balancer
	var conf = "random"  //指定默认的配置算法
	if len(os.Args) > 1 {  //os.Args获取命令行参数个数
		conf = os.Args[1] //取第一个参数
	}
	//判断输入的参数
	if conf == "random" {
		balancer = &balance.RandomBalance{} //传地址，因为*RandomBalance实现了DoBalance方法
		fmt.Println("use random balancer!")
	}else if conf == "roundrobin" {
		balancer = &balance.RoundrobinBalance{}
		fmt.Println("use roundrobin balancer!")
	}
	//执行具体的负载均衡策略
	for {
		inst,err := balancer.DoBalance(insts)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}