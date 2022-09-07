package main

import (
	"fmt"
)

func updateName(x string) string{
	x = "wedge"
	return x
}
func updateMenu( o map[string]float64){
	o["coffee"] = 2.99
}
func main(){
	names:= "tofu"
names = updateName(names)
fmt.Println(names)
	fmt.Println("MAPS IN GOLANG !!")
	menu:= map[string] float64{
   "soup" : 4.99,
   "burger": 6.99,
   "noodles" : 5.99,
   "vadapav": 5.05,
	}
	fmt.Println(menu)	
fmt.Println(menu["vadapav"])
 updateMenu(menu)
fmt.Println(menu)	
for key,value:= range menu{
	fmt.Printf("key : %v,value :%v\n",key,value )
}
// int as key
phoneBook := map[int]string{
	2479687:"rohini",
	1905456: "pushpa",
	8951569: "nandu",
}
fmt.Println(phoneBook)
fmt.Println(phoneBook[2479687])
// to update a key-value pair 
phoneBook[2479687] = "chandu"
fmt.Println(phoneBook)
fmt.Println(phoneBook[2479687])
}

