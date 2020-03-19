package main

import (
	"fmt"
	// "github.com/xpace/playground/stringutil"
	// "time"
	"github.com/xxpace/playground/web"
	// "strings"
	// "os"
	// "io"
)

type User struct{
	name string
	age int8
}

func print(content interface{}){
	fmt.Println(content);
}

func printUser(user User){
	print(user.name)
	print(user.age)
}

func testSlice(){
	// arr := [] int {1,2,3,4}
	var numbers []int
	numbers = append(numbers,3,4,5)
	for index,num :=range "go"{
		print(index);
		print(num);
	}
}

func testMap(){
	tabelMap := make(map[string]int)
	tabelMap["xuhe"] = 1
	tabelMap["zhaoyang"] = 2
	delete(tabelMap,"xuhe")
	age,ok :=tabelMap["xuhe"]
	if(ok){
		print(age);
	}else{
		print("not existed")
	}

	print(string(1))
}
type Animate interface{
	say()
}

type Dog struct{

}
func (dog Dog) say(){
	fmt.Println("dog :wangwang ")
}

type Checken struct{

}

func (Checken) dodo(){
	fmt.Printf("checken:dodo")
}

func (checken Checken) say(){
	fmt.Printf("checken:gugu")
	checken.dodo()
}

func testInterface(){
	var animate Animate

	animate = new(Dog)
	animate.say()

	animate = new(Checken)
	animate.say()
}

func sum(s []int, c chan int){
	sum := 0
	for _,v := range s{
		sum += v
	}
	c <- sum
}

func main(){
	// admin := User{"xuhe",30}
	// var ptr *User = &admin
	// printUser(*ptr);
	// testSlice();
	// testMap();
	// testInterface();

	// s := []int {1,2,3,4,5,6}
	// c := make(chan int)
	// half := len(s)/2
	// go sum(s[:half],c)
	// go sum(s[half:],c)
	// x,y := <-c,<-c
	// fmt.Println(x,y,x+y)

	web.StartServer();

	// reader := strings.NewReader("hello world")
	// p := make([]byte,4)

	// for{
	// 	n,err := reader.Read(p)
	// 	if(err!=nil){
	// 		if(err==io.EOF){
	// 			fmt.Println("EOF:",n)
	// 			break;
	// 		}
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println(n,string(p[:n]))
	// }
}