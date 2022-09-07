package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bill struct {
	name string
	menu map[string]float64
	tip  float64 
}

func newBill(name string) Bill {
	b:= Bill{
		name: name,
		menu: map[string]float64 {"coffee": 2.99,"water" : 4.99,},
		tip: 0,
	}
	return b
}
func (b *Bill) updateTip(tip float64){
	b.tip = tip
}
func (b *Bill) updateMenu(name string,price float64) {
	b.menu[name] = price
}

func (b *Bill) format()string{
fs := "Bill Breakdown : \n "



var total float64 = 0 
for k,v:= range b.menu{
	fs += fmt.Sprintf("%-25v ...%v\n", k+":",v)
	total += v
}
 
fs += fmt.Sprintf("%v ...$%0.2f","total: ", total+b.tip)
return fs
}

func main() {
	customerBill := newBill("lexi")
	customerBill.updateTip(6)
	customerBill.updateMenu("soup",5.99)
	customerBill.updateMenu("sambar",3.99)
	fmt.Println(customerBill.format())
	
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Println("Type your name : ")
	input:= scanner.Text()
	fmt.Printf("Hello %v !! \n",input)
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your birth year : ")
	age,_:= reader.ReadString('\n')
    myyear,_:= strconv.ParseInt(strings.TrimSpace(age),10,64)
fmt.Println(2022-myyear)

}
