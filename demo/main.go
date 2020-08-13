package main

import "fmt"

type i1 interface {
	 i1_fun()
}

type i1_struct struct {
	a int
}

func(i1_struct) i1_fun(){

}
func main() {
	var a i1
	var as = i1_struct{10}
	a = as
	aa,ok := a.(i1_struct)
	fmt.Println(aa,ok)

	switch v:=a.(type) {
	case i1_struct:
	  fmt.Println(v,"i1struct")
	}
}
