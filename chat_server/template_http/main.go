package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var myTemplate *template.Template

type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name  string
	Title string
	Age   int
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	var arr []Person
	p1 := Person{Name: "Ponny", Age: 10, Title: "Ponny的个人网站"}
	p2 := Person{Name: "Jack", Age: 12, Title: "Jack的个人网站"}
	p3 := Person{Name: "Robin", Age: 14, Title: "Robin的个人网站"}
	arr = append(arr, p1)
	arr = append(arr, p2)
	arr = append(arr, p3)

	resultWriter := &Result{}
	io.WriteString(resultWriter, "hello world")
	err := myTemplate.Execute(w, arr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("template render data:", resultWriter.output)
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("./index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:5000", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
