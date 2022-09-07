package main

import (
	"fmt"
	"strings"
)

func IndiaMsg(n string) {

	fmt.Printf("I hope you like Indian %v!!\n",n)	
}
func IndiaMotto(){
	fmt.Println("Namste Turkey !! We believe Unity in diversity.")
}
func TurkeyMsg(n string){
	fmt.Printf("I hope you like  Turkish to %v !!\n",n)
	
}
func ProAdder(values ...int)(int,string){
	total:= 0
	for _,value:= range values{
     total += value
	}
	return total,"Im tired man."
}
func PeaceMsg(n []string, f func(string)){
	for _,v:= range n{
		f(v)
	}
}

func names(n string)(string,string){
	caps:= strings.ToUpper(n)
word:= strings.Split(caps, " ")
fmt.Println(word)
var FirstNames []string

for _,val:= range word{
	FirstNames = append(FirstNames,val[0:])
}
if len(FirstNames)>1 {
	return FirstNames[0], FirstNames[1]
}
	return FirstNames[0], "_"
}
 
func main(){
	first,second := names("Rohini Munnangi")
	fmt.Println(first,second)
IndiaMotto()
//TurkeyMsg("India")
PeaceMsg( []string {"Movies", "Food", "cutlure"}, IndiaMsg)
PeaceMsg( []string {"Movies", "Food", "cutlure"}, TurkeyMsg)
Proresult,bye := ProAdder(5,7,9,100)
fmt.Printf("final value is %v,%v",Proresult, bye)
}

