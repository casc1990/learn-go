package main

import (
	"bufio"
	"os"
	"fmt"
)
//输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数
func count(str string) (worldcount, spacecount, numbercount, othercount int) {
	s := []rune(str)
	for _,v := range s {
		switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':   //大小写是一种情况
				worldcount++
			case v == ' ':
				spacecount++
			case v >= '0' && v <= '9':
				numbercount++
			default:
				othercount++
			
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:",err)
		return
	}
	wc,sc,nc,oc := count(string(result))
	fmt.Printf(" world count:%d\n space count:%d\n number count:%d\n other count:%d",wc,sc,nc,oc)

}
