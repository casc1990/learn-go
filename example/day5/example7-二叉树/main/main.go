package main

import (
	"math/rand"
	"fmt"
)

//二叉树：如果每个节点有两个指针分别用来指向左子树和右子树，我们把这样的结构叫做二叉树

type Student struct {
	Name string
	Age int
	Score float32
	left *Student  //指向左子树的指针
	right *Student   //指向右子树的指针

}

//生成二叉树
func Btree()  *Student {
	//根节点
	var root *Student = &Student{
		Name: "root",
		Age: 18,
		Score: rand.Float32() * 100,
	}

	//左叶子根节点
	var left1 *Student = &Student{
		Name: "root-left1",
		Age: 18,
		Score: rand.Float32() * 100,
	}
	//左叶子根节点加入根节点
	root.left = left1 //加入根节点
	//左叶子根节点的子节点
	var left2 *Student = &Student{
		Name: "root-left2",
		Age: 18,
		Score: rand.Float32() * 100,
	}
	//左叶子节点加入左叶子根节点
	left1.left = left2 //加入左叶子根节点

	//右叶子根节点
	var right1 *Student = &Student{
		Name: "root-right1",
		Age: 18,
		Score: rand.Float32() * 100,
	}
	//又叶子根节点加入根节点
	root.right = right1

	return root
}

//遍历二叉树(使用递归)
func trans(root *Student) {
	//递归的退出条件是：被递归的节点没节点了
	if root == nil {
		return
	}
	//打印节点内容
	//fmt.Println(*root)

	//使用递归，不断的打印每一层节点的left子节点(比如：第一次是root-left1，第二层就是root-left1.left...)
	trans(root.left)
	trans(root.right)
}


func main() {
	root := Btree()
	trans(root)
}