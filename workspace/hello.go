package main

import "fmt"

func main(){
	fmt.Println("hello world")

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
}