package main
//goroutine说到底其实就是线程，goroutine是通过Go的runtime管理的一个线程管理器.goroutine通过go关键字实现
import (
	"fmt"
	"runtime"
	"strconv"
)

func say(s string) {
	for i := 0;i < 5; i++ {
		//runtime.GOMAXPROCS(n)  同时使用几个线程
		runtime.Gosched()  //表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}

//channel--通道(默认情况下，channel接收和发送数据都是阻塞的)
func sum(a []int, c chan int) {
	total := 0
	for _,v := range a {
		total += v
	}
	c <- total  //发送total的值到chan
}

//channel的range和close
func fibonacci (n int, c chan int) {
	x, y := 1,1
	for i :=0;i < n; i++ {
		fmt.Printf("第%d次:\n",i)
		c <- x  //向channel中发送数据，同时能存储多少，取决于channel的Buffered
		x, y = y, x+y
	}
	close(c)  //关闭channel，关闭了就不能读取和接收了
}

//channel的select
func fibonacci1(c,quit chan int) {
	x ,y := 1,1
	for {
		select {
		case c <- x:
			x , y = y , x + y
		case <- quit:
			fmt.Println("quit..")
			return
		}
	}
}

func main() {
	go say("hello world")  //通过go关键字，开一个新的goroutines执行
	say("hello mingri")  //当前goroutines执行

	//channels 通道--必须使用make创建channel
	//定义接受int类型的channel
	a := []int{7,2,8,-9,4,0}
	cl := make(chan int)  //定义channel时，也需要指定发送或者接受着的值的类型
	cs := make(chan string) //定义接受string类型的channel
	cf := make(chan interface{})
	fmt.Println(cl, cs, cf)  //0xc042084000 0xc042084060 0xc0420840c0
	go sum(a[:len(a)/2],cl) //通过go关键字，开一个新的goroutines执行
	go sum(a[len(a)/2:],cl)
	x := <- cl  //从chan中接收数据，并赋值给x
	y := <- cl //接收第二个数据，没有数据了，在接收就报错了
	fmt.Println("x:"+strconv.Itoa(x)+" y:"+strconv.Itoa(y)+" x+y:"+ strconv.Itoa(x+y)) //x:-5 y:17 x+y:12

	////channel的range和close
	c := make(chan int,10)  //指定channel最多可以同时接收或者发送10个
	go fibonacci(20,c)  //cap(c)等于10
	for i := range c {
		fmt.Println("---:",i)
		fmt.Println(i)  //打印channel里的数据（取数据）
	}

	//channel的select
	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i :=0;i < 10; i++ {
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	fibonacci1(c1,quit)
}