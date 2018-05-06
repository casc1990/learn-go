package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`   //指定打包或者序列化时的字段名称
	NickName string `json:"nickname"`
	Age  int
	Birthday string
	Sex string
	Email string
	Phone string
}

func testStruct() {
	user1 := User{
		UserName: "user1",
		NickName: "哈哈",
		Age: 	28,
		Birthday: "2018/08/08",
		Sex: "男",
		Email: "mahuatong@qq.com",
		Phone: "110",
	}
	data,err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal failed,err:",err)
		return
	}
	fmt.Println(string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "man"

	data,err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal failed,err:",err)
		return
	}
	fmt.Println(string(data))

}


func testSlice() {
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	for i:=1;i<5;i++ {
		m["username"] = fmt.Sprintf("user%d",i)
		m["age"] = fmt.Sprintf("%d",i+20)
		if i%2==1 {
			m["sex"] = "man"
		}else {
			m["sex"] = "woman"
		}
		s = append(s,m)
	}
	data,err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed,err:",err)
		return
	}
	fmt.Println(string(data))

}

func main() {
	testStruct() //{"username":"user1","nickname":"哈哈","Age":28,"Birthday":"2018/08/08","Sex":"男","Email":"mahuatong@qq.com","Phone":"110"}
	testMap() //{"age":29,"sex":"man","username":"user2"}
	testSlice()


}

