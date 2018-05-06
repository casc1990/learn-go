package main

import "fmt"

//var ch <-chan int:只读channel
//var ch chan<- int:只写channel
//var ch chan int:默认channel（读写）

func send(ch chan<- int,exitChan chan struct{}) {
	for i:=0;i<10;i++ {
		ch <- i
	}
	close(ch)  //数据发送完了，就关闭channel，别人就可以检测到channel关闭了，然后读取完退出了

	//goroute执行完，发送信号
	var exitStuct struct{}
	exitChan <- exitStuct

}

func recv(ch <-chan int,exitChan chan struct{}) {
	for {
		//判断channel是否关闭
		v,ok := <- ch
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(v)
	}

	//goroute执行完，发送信号
	var exitStuct struct{}
	exitChan <- exitStuct
}

func main() {
	ch := make(chan int,10)
	exitChan := make(chan struct{},2)
	go send(ch,exitChan)
	go recv(ch,exitChan)

	//判断goroute是否执行完（读取exitChan元素）
	var total = 0
	for _ = range exitChan { //循环读取数据，如果没读到数据会阻塞
		total++
		if total == 2 {  //一共读到2个元素，说明goroute执行完了，然后退出
			break
		}
	}
}