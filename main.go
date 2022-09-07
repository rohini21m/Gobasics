package main

import (
	"fmt"
)

func main(){
	fmt.Println("I am Rohini Munnangi. I am a soul")
	var name = "timothee"
	var namethree string 
	namethree = "pushpa"
	 leader  := "Rohini"
	 var nametwo string = "rosie"
	 var age int
	 age = 65 
	 ageone:= 100
	 fmt.Println(name, leader,nametwo, namethree)
	 fmt.Printf("nobody wants to be  %v today\n", ageone)
	 fmt.Printf("all struggle to live upto %v happily !!\n",age)
	 var lakecharles  int32 = 6700
	 var rollnum uint16 = 1000
     var citytemp float32 = 66.7
	 fmt.Printf("population of my city is %v but not %v \n",lakecharles,rollnum )
	 fmt.Printf("temperatire of city is %v deg celsius ",citytemp)
	 century:= "21"
	 country := "India"
// 
	 fmt.Printf("Welcome to %q. This is my %q\n", century,country) 
	 fmt.Printf("My total score is %0.1f\n", 334.222)

// sprinft savers the string value in a variable.
var rohini = fmt.Sprintf("Welcome to %q. This is my %q\n", century,country)
fmt.Println("saved info :",rohini)

ages := [4]int{2,8,19,900}
names:= [3] string{"chandu", "nandu" ,"sai"}
fmt.Println(ages,len(ages), names,len(names))
cities := []string{"vijayawada","Bombay"}


cities = append(cities, "Goa") // this updates var cities that has been declared before.
fmt.Println(cities)
cities[2] = "Delhi"
fmt.Println(cities)
rangeOne:= cities[0:2]
rangeOne= append(rangeOne, "Hyderabad")
rangeTwo:= cities[1:]
rangeThree:= cities[:3]
fmt.Println(cities,len(cities),rangeOne,rangeTwo,rangeThree)



}
