package main

import (
	"sync"
	"time"
	"fmt"
)

var (
	m  = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i:=1;i<t.n;i++ {
		sum *= uint64(i)
	}

	//因为同一时间只能有一个协程写文件，所以使用互斥锁
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i:=0;i<50;i++ {
		t := &task{n:i}
		go calc(t)
	}

	time.Sleep(time.Second*10)
	lock.Lock()
	for k,v := range m {
		fmt.Printf("%d阶乘=%v\n",k,v)
	}
	lock.Unlock()
}