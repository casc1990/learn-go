package main

import "fmt"

//接收3个channel参数
func calc(taskChan chan int,resChan chan int,exitChan chan bool) {
	//从taskChan取值
	for v := range taskChan {
		flag := true
		//判断是否是素数（只能被本身和1整除）
		for i:=2;i<v;i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		//是素数，将结果写入到resChan中
		if flag {
			resChan <- v
		}
	}

	//goroute执行完，退出，发送一个true给exitChan
	fmt.Println("exit")
	exitChan <- true
}


func main() {
	intChan := make(chan int,1000)
	resultChan := make(chan int,1000)
	exitChan := make(chan bool,8)

	//向intChan写数据
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i:=0;i<8;i++ {
		go calc(intChan,resultChan,exitChan)
	}

	//等待所有计算的goroute全部退出
	go func(){
		for i:=0;i<8;i++ {
			<- exitChan
			fmt.Println("goroute ",i,"exited!")
		}
		//exitChan接收到8个退出，说明任务已经做完了，退出resultChan
		close(resultChan)
	}()

	//因为resultChan已经退出，取到最后一个元素，for循环也就退出了
	for v:= range resultChan {
		fmt.Println(v)
	}
}