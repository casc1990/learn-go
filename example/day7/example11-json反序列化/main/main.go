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

func testStruct() (ret string,err error) {
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
		err = fmt.Errorf("json marshal failed,err:",err)
		return
	}
	ret = string(data)
	return
}

func testMap() (ret string,err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "man"

	data,err := json.Marshal(m)
	if err != nil {
		err = fmt.Errorf("json marshal failed,err:",err)
		return
	}
	ret = string(data)
	return
}

func UnmarshalStruct() {
	data,err := testStruct()
	if err != nil {
		fmt.Println("json marshal failed err:",err)
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data),&user1)
	if err != nil {
		fmt.Println("Unmarshal struct failed!,err:",err)
		return
	}
	fmt.Println(user1)
}


func UnmarshalMap() {
	data,err := testMap()
	if err != nil {
		fmt.Println("json marshal failed err:",err)
		return
	}
	var m map[string]interface{}
	m = make(map[string]interface{})
	err = json.Unmarshal([]byte(data),&m)
	if err != nil {
		fmt.Println("Unmarshal struct failed!,err:",err)
		return
	}
	fmt.Println(m)
}


func main() {
	UnmarshalStruct() //{user1 哈哈 28 2018/08/08 男 mahuatong@qq.com 110}
	UnmarshalMap() //map[sex:man username:user2 age:29]
}

