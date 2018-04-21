package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
	"sync/atomic"
)


//互斥锁:同一时间只能有一个线程操作，定义格式： var mu sync.Mutex
//读写锁：一个读写锁同时只能有一个写者或多个读者（与CPU数相关），但不能同时既有读者又有写者（适用于读多写少）
//原子操作：同一时刻数字是相同的 atomic.AddInt32:保证写数字是原子操作的。atomic.LoadInt32：保证读数据是原子操作的

var lock sync.Mutex
var rwlock sync.RWMutex

//互斥锁
func testMutex() {
	a := make(map[int]int)
	a[8] = 8
	a[3] = 3
	a[2] = 2
	a[5] = 5

	for i :=0;i <3;i++ {
		go func(b map[int]int) {
			//两个goroute都操作字典，所以要加互斥锁
			lock.Lock() //上锁
			b[8] = rand.Intn(100)
			lock.Unlock() //释放锁
		}(a)
	}

	//因为goroute在后台执行，程序执行到读map操作时，发现有线程在写，所以也要加锁
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second) //等待goroute执行完

	}

//读写锁
func testRLock() {
	var count int32
	a := make(map[int]int)
	a[8] = 8
	a[3] = 3
	a[2] = 2
	a[5] = 5

	//模拟写请求，使用读写锁中的写锁
	for i :=0;i < 3;i++ {
		go func(b map[int]int) {
			rwlock.Lock() //上写锁
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond) // 模拟处理逻辑
			rwlock.Unlock() //释放写锁
		}(a)
	}

	//模拟读请求。读锁不影响其他线程的读
	for i :=0;i < 100;i++ {
		go func(b map[int]int){
			for {
				rwlock.RLock()  //上读锁
				time.Sleep(time.Millisecond) //模拟业务处理逻辑
				rwlock.RUnlock() //释放锁
				atomic.AddInt32(&count,1) //atomic保证原子操作，处理次数++
			}
		}(a)
	}

	time.Sleep(time.Second * 3)  //等待goroute执行完
	fmt.Println(atomic.LoadInt32(&count))  //打印处理次数 atomic.LoadInt32(&count)保证读取的原子性
}

func main() {
	testMutex()
	testRLock()
}

