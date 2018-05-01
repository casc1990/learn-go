package main

//负载均衡算法的高层封装

import (
	"learn-go/example/day7/example1-负载均衡/balance"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	var balanceName = "random"  //指定默认的配置算法
	if len(os.Args) > 1 {  //os.Args获取命令行参数个数
		balanceName = os.Args[1] //取第一个参数
	}
	//调用DoBalance函数，DoBalance函数封装了底层的DoBalance方法
	for {
		inst,err := balance.DoBalance(balanceName,insts)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}


