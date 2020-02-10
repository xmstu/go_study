package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testStruct() (ret string, err error) {
	user1 := &User{
		UserName: "user1",
		NickName: "黄片哥",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
	ret = string(data)
	return
}

func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func testMap() (ret string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
	ret = string(data)
	return

}

func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s = append(s, m)
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 29
	m["sex"] = "female"
	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed err:", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func test2() {
	data, err := testMap()
	if err != nil {
		fmt.Println("test map failed, ", err)
		return
	}

	var m map[string]interface{}
	// 反序列化
	err = json.Unmarshal([]byte(data), &m)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(m)
}

func test() {
	data, err := testStruct()
	if err != nil {
		fmt.Println("test struct failed, ", err)
		return
	}

	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(user1)
}

func main() {
	//	testStruct()
	//	testInt()
	//	testMap()
	//	testSlice()

	test()
	test2()
}
