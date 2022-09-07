package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	GratitudeMsg := "I am thankful and open to the abundance,love universe has for me."
	fmt.Println(strings.ToUpper(GratitudeMsg))
	fmt.Println(strings.IndexAny(GratitudeMsg,"open"))
	fmt.Println(strings.Index(GratitudeMsg,"open"))
	fmt.Println(strings.Contains(GratitudeMsg, "universe"))
	fmt.Println(strings.ContainsAny(GratitudeMsg,"abundance"))
	fmt.Println(strings.ReplaceAll(GratitudeMsg,"abundance","propserity"))
	fmt.Println("original msg :",GratitudeMsg)
	fmt.Println(strings.Split(GratitudeMsg,"open"))

	languages:= []string{"telugu","hindi","malayalam","urdu","turkish"}
	country:= []int{20,10,90,76,555}
	sort.Ints(country)
	fmt.Println(country)
fmt.Println(sort.SearchInts(country,76))
	sort.Strings(languages)
	fmt.Println(languages)
	fmt.Println(sort.SearchStrings(languages,"urdu"))
	for index,value := range languages{
	fmt.Printf("The value at index %v is %v\n",index,value)	
	value = "new string"
	}
	fmt.Println(languages)
}