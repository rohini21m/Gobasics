package main
import (
	"io/ioutil"
	"strconv"
	"testing"
	"net/http"
	"net/http/httptest"
)
func TestDoubleHandler(t *testing.T){
	tt:= []struct{
		name string 
		value string 
		status int 
		double int
		err string
	}{
		{name: "double of three",value:"3",double:12},
		{name: "missing value ",value:"",err:"missing value"},
		{name: "not a number",value:"oops",err:"not a number: oops"},
	}
	for _,tc := range tt{
	t.Run(tc.name,func(t *testing.T){
req, err:= http.NewRequest("GET","localhost://8080/double?v="+tc.value,nil)
if err!=nil{
	t.Fatalf("could not create request : %v",err)
}
rec:= httptest.NewRecorder()
 doubleHandler(rec,req)
 res:= rec.Result()
 
 defer res.Body.Close()
 b, err:= ioutil.ReadAll(res.Body)
 if err!=nil{
	t.Fatalf("could not read response, got : %v", err)
 }
 if  tc.err!=""{
	// d something
	if res.StatusCode!= http.StatusBadRequest{
		t.Errorf("expected value bad request; but got  %v",res.StatusCode)
	}
	if msg := string(b);msg!= tc.err{
		t.Errorf("expected message %q; got %q ",tc.err,msg)
	}
	return 
 }
 if res.StatusCode!= http.StatusOK{
	t.Errorf("not expected Status-OK, got : %v",res.Status)
 }
 
 // parse content of response
 d,err:= strconv.Atoi(string(b))
 if err!=nil{
	t.Fatalf("expected an integer; got: %s",b)
 }
 if d!= 3{
	t.Fatalf("expected double to be 12 but got : %v",d)
 }

})
}

// tetsing http routing 
func TestRouting(t testing*T){
	srv:= httptest.NewServer(handler())
	defer srv.Close()
	res, err:= http.Get(fmt.Sprintf("%s/double?v=3",srv.URL))
	if err!=nil{
		t.Fatalf("could not send get request: %v",err)
	}
	b, err:= ioutil.ReadAll(res.Body)
 if err!=nil{
	t.Fatalf("could not read response, got : %v", err)
 }
 if  tc.err!=""{
	// d something
	if res.StatusCode!= http.StatusBadRequest{
		t.Errorf("expected value bad request; but got  %v",res.StatusCode)
	}
	if msg := string(b);msg!= tc.err{
		t.Errorf("expected message %q; got %q ",tc.err,msg)
	}
	return 
 }
 if res.StatusCode!= http.StatusOK{
	t.Errorf("not expected Status-OK, got : %v",res.Status)
 }
 
 // parse content of response
 d,err:= strconv.Atoi(string(b))
 if err!=nil{
	t.Fatalf("expected an integer; got: %s",b)
 }
 if d!= 3{
	t.Fatalf("expected double to be 12 but got : %v",d)
 }
}

