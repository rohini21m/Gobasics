package main

import "fmt"
func updateName(x *string){
	*x = "wedge"
}
func main(){
	name:= "football"
	m:= &name
	fmt.Println(name,m,*m)
	fmt.Println("Lets look at pointers in Go. ")

	rohini := "chandu"
	fmt.Printf("memory address of rohini :%v\n",&rohini)
	p:= &rohini
	fmt.Printf("value stored in p: %v and its memory address: %v\n",*p,&p)
	*p = "Munnangi"
	fmt.Println(rohini)
  updateName(m)
  fmt.Println(name)
}

