package main

import "fmt"

//字符串翻转函数
func reverse(s string) string {
	var result string
	lens := len(s)
	for i := 0;i < lens;i++ {
		result = result + fmt.Sprintf("%c",s[lens-i-1]) //通过下标从后往前取(下标等于len-1)
	}
	return result
}

//通过slice实现字符串翻转
func reverse1(s string) string {
	var result []byte
	tmp := []byte(s)
	lens := len(s)
	for i :=0;i < lens;i++ {
		result = append(result,tmp[lens-i-1])
	}
	return string(result)
}

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + " " + str2 //字符串拼接
	fmt.Print(str3,"\n")
	fmt.Printf("%s %s\n",str1,str2)

	str4 := fmt.Sprintf("%s %s\n",str1,str2)
	l := len(str4)
	fmt.Printf("len(str4)=%d\n",l)
	substr := str4[0:5]  //字符串切片
	fmt.Print(substr,"\n")
	substr = str4[6:]
	fmt.Print(substr,"\n")

	result := reverse(str4)
	fmt.Print(result,"\n")  //dlrow olleh

	result1 := reverse1(result)
	fmt.Print(result1) //hello world
}
