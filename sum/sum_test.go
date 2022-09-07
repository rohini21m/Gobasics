package sum

import(
     "testing"
)

func TestInts(t *testing.T){
	//t.Fail()
	//t.Errorf(" test failed to return sum of integers")
	//t.Fatalf("test failed and stp running")// t.Fatalf will immediately stop the function exexution and dont 
	// proceed to next line wherre ad t.Errorf will continue to execute the func
tt:= []struct{
	name string
	numbers  []int
	sum int
}{
	{"one to five",[]int{11,22,33,44,55},165},
	{"1to -1",[]int{1,-1, 2, -2},0},
	{"no numbers",nil,0},
}
for _,tc:= range tt{
s := Ints(tc.numbers...)
	if s!=tc.sum{
		t.Errorf("sum of  %v; should be ; %v,got %v", tc.name,tc.sum,s)
	}

}
// SubTests
// for _,tc:= range tt{
//	t.Run(tc.name,func(t *testing.T){
// 	s := Ints(tc.numbers...)
// 		if s!=tc.sum{
// 			t.Fatalf("sum of  %v; should be ; %v,got %v", tc.name,tc.sum,s) // only particular 
// test is stopped if failed not the rest 
// 		}
//})
// 	}
	// s:= Ints(1,2,3,4,5)
	// if s!=15{
	// 	t.Errorf("result should be  15 but got %v",s)
	// }
}

func TestVoo(t *testing.T){
t.Errorf("test Voo is failed always")
}