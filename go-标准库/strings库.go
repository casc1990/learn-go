package main

import (
	"fmt"
	"strings"
	"strconv"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url,"http")  //字符串以什么开头
	if !result {
		url = fmt.Sprintf("http://%s",url)
	}
	result1 := strings.HasSuffix(url,"/")  //字符串以什么结尾
	if !result1 {
		url = fmt.Sprintf("%s%s",url,"/")
	}
	return url
}


func main() {
	var (
		url string = "1.1.1.1/user/login"
		str string
	)
	//fmt.Scanf("%s",&url)
	url = urlProcess(url)
	fmt.Printf("%s\n",url) //输出:http://1.1.1.1/user/login/

	str = " hello World! i Love you, a,b,c 2 t A \n"
	//strings.Index(s string, str string) int：判断str在s中首次出现的位置,如果没有出现，返回-1
	fmt.Println("o在字符串中首次出现的位置是:",strings.Index(str,"o")) //5
	fmt.Println("o在字符串中最后出现的位置是:",strings.LastIndex(str,"o")) //22

	str1 := strings.Replace(str,"you","me",1)
	fmt.Println("字符串替换后的内容:",str1) //hello World! i love me, a,b,c 2 t A
	str2 := strings.Replace(str,"o","e",-1) //-1全部替换
	fmt.Println("字符串全部替换后的内容:",str2)

	str3 := strings.Count(str,"o")
	fmt.Println("o在字符串中出现:",str3) //字符出现的次数

	str4 := strings.Repeat("hello ",4)
	fmt.Println("字符串重复4次后是:",str4) //字符串重复4次后是: hello hello hello hello

	str5 := strings.ToLower(str)
	fmt.Println("字符串全部转换成小写后是:",str5) //转小写

	str6 := strings.ToUpper(str)
	fmt.Println("字符串全部转换成大写后是:",str6) //转大写

	str7 := strings.TrimSpace(str)
	fmt.Println("去除行首和行位的空白字符是:",str7)
	str8 := strings.Trim(str,"he")
	fmt.Println("去除行首和行位的指定字符是:",str8)
	fmt.Println("去除行首指定字符是:",strings.TrimLeft(str,"he"))
	fmt.Println("去除行尾指定字符是:",strings.TrimRight(str,"\n\r"))

	str9 := strings.Fields(str)
	fmt.Println("以空格分割，将字符串转换成slice:",str9) //[hello World! i Love you, a,b,c 2 t A]
	fmt.Println("第六个元素:",str9[5]) //a,b,c
	//指定分割符
	str10 := strings.Split(str,",")
	fmt.Println("以,分割，第1个元素是:",str10[0]) //hello World! i Love you

	//slice到str的逆向转换
	str11 := strings.Join(str9," ")
	fmt.Println("以,分割，将slice转换成str为:",str11)

	//int转换成str
	str12 := strconv.Itoa(100)
	fmt.Printf("将int转换成%T",str12)

	//str转换为int
	str13,err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("不能转换...")
	}
	fmt.Println("将str转换成int：",str13)

}
