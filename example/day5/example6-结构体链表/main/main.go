package main

import (
	"fmt"
	"math/rand"
)

//链表：每个节点包含下一个节点的地址，这样把所有的节点串起来了，通常把链表中的第一个节点叫做链表头

type Student struct {
	Name string
	Age int
	Score float32
	next *Student
}

//遍历结构体链表，并打印结构体内容
func trans(p *Student) {
	for p != nil {  //首先循环链表头，链表头肯定有next，然后在判断其他链表是否有next方法
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

//尾部插入链表
func insertTail(p *Student) {
	//循环加入链表（思路：记录最后链路的位置，将创建的结构体加入）
	var tail = p //tail报存最后的结构体位置（默认从第一个位置开始）
	for i :=0;i<10;i++ {
		var str = Student{  //创建结构体
			Name: fmt.Sprintf("sut%d",i),
			Age: rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &str //加入结构体
		tail = &str //移动最后结构体的位置
	}
}

//从头部开始插入链表
func insertHead(p **Student) {
	for i:=0;i<10;i++ {
		var str = Student{  //创建结构体
			Name: fmt.Sprintf("sut%d",i),
			Age: rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		str.next = *p
		*p = &str
	}
}

//删除节点链表中的某个节点
func delnode(p *Student) {
	//定于临时变量，报存被删除节点的上一个节点
	var prev *Student = p

	//循环链表
	for p != nil {
		//匹配要删除的结构体
		if p.Name == "sut6" {
			//将要删除的结构体的上一个结构体的next指向被删除结构体的下一个结构体
			prev.next = p.next
			break
		}
		//如果没找到，向后调位继续匹配
		prev = p
		p = p.next
	}
}


//增加链表中的节点
func addnode(p *Student,newNode *Student) {
	for p != nil {
		//在stu5后面插入，将新节点的next指向要插入节点的next，将要插入节点的next指向新节点的next
		if p.Name == "sut7" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		//如果没找到，向后调位继续匹配
		p = p.next
	}
}


func main() {
	var head  *Student = &Student{
		Name: "lala",
		Age: 20,  //未赋值的字段默认为空
	}

	var str1 = Student{
		Name: "hehe",
		Age: 28,
	}

	head.next = &str1  //把str1的内存地址加入结构体链表中
	trans(head) //遍历链表

	//从尾部插入链表
	//insertTail(head)
	//trans(head) //遍历链表

	//从头部插入链表
	insertHead(&head) //如果想要改变指针变量的值，需要传指针的指针
	trans(head) //遍历链表

	//删除链表中某个节点
	delnode(head)
	trans(head) //遍历链表

	//增加链表中的节点
	var str2 = Student{
		Name: "sut6",
		Age: 28,
	}
	addnode(head,&str2)
	trans(head) //遍历链表
}

