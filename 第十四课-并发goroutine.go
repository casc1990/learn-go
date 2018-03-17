package main
//goroutine说到底其实就是线程，goroutine是通过Go的runtime管理的一个线程管理器.goroutine通过go关键字实现
import (
	"fmt"
	"runtime"
	"strconv"
	"time"
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
	close(c)  //关闭channel，关闭的信道会禁止数据流入, 是只读的
}

//channel的select（select和switch一样,可处理一个或多个channel的发送与接收 ）
func fibonacci1(c,quit chan int) {
	x ,y := 1,1
	for {     //因为select执行完就会退出，使用for无限接收或读取数据
		select {
		case c <- x:  //根据case处理多个或者同一个channel
			x , y = y , x + y
		case <- quit:  //如果quit这个channel收到数据，程序退出
			fmt.Println("quit..")
			return
		default:  //select的default语句
			fmt.Println("done!")
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
	var cf chan int = make(chan int)  //也可以这样定义channel
	fmt.Println(cl, cs, cf)  //0xc042084000 0xc042084060 0xc0420840c0
	go sum(a[:len(a)/2],cl) //通过go关键字，开一个新的goroutines执行
	go sum(a[len(a)/2:],cl)
	x := <- cl  //从chan中接收数据，并赋值给x
	y := <- cl //接收第二个数据，没有数据了，在接收就报错了
	fmt.Println("x:"+strconv.Itoa(x)+" y:"+strconv.Itoa(y)+" x+y:"+ strconv.Itoa(x+y)) //x:-5 y:17 x+y:12

	////channel的range和close
	c := make(chan int,10)  //指定channel最多可以同时接收或者发送10个
	go fibonacci(20,c)  //cap(c)等于10
	for i := range c {    //range不等到信道关闭是不会结束读取的,range会阻塞当前goroutine,产生死锁
		fmt.Println("---:",i)
		fmt.Println(i)  //打印channel里的数据（取数据）
	}

	//channel的select
	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i :=0;i < 10; i++ {
			fmt.Println(<- c1)  //接收channel数据
		}
		quit <- 0  //想quit的channel发送数据
	}()
	fibonacci1(c1,quit)

	//select的超时设置
	c4 := make(chan  int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c4:
				fmt.Println(v)
			case <- time.After(3 * time.Second):
				fmt.Println("timeout..")
				o <- true
				break
			}
		}
	}()
	<- o
}