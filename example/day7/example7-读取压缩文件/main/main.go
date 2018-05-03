package main

import (
	"os"
	"fmt"
	"compress/gzip"
	"bufio"
)

func main() {
	FName := "MyFile.gz"
	fi,err := os.Open(FName)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v,can't open %s: error:%s\n",os.Args[0],FName,err)
		os.Exit(1)
	}
	defer fi.Close()

	fz,err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr,"open gzip failed,err: %v\n",err)
		return
	}
	r := bufio.NewReader(fz)
	for {
		line,err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
