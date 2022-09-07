package main

import "fmt"

func main(){

	fmt.Println("Working with Boolean functions")
	age := 99 
	fmt.Println(age<=25)
	fmt.Println(age>=99)
	fmt.Println(age==99)
	fmt.Println(age<100)

if age>= 10{
	fmt.Println("Still not applicable untill 99")
}else if age <= 100{
fmt.Println("Appropriate Age till 99")
} else {
	fmt.Println("ITS a kid show")
}
languages:= []string{"telugu","hindi","malayalam","urdu","turkish"}
for index,value:= range languages{
	if index == 1 {
		fmt.Printf("continuing the position at %v the value is %v \n",index, value)
		continue
	}
	if index>2{
		fmt.Println("breaking the cycle at position",index)
		break
	}
	fmt.Printf("value at position %v is %v \n",index, value)
	
}

}