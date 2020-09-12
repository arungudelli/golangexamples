package main

import (
    "strconv" 
    "fmt"
)
func main() {
    var int64Value int64 = 100
    var stringValue string
	// stringValue = strconv.Itoa(intValue)
	stringValue = strconv.FormatInt(int64Value, 10)
    fmt.Printf("%T, %v\n",int64Value,int64Value)
	fmt.Printf("%T, %v\n",stringValue,stringValue)
	
	var n int64 = 100
	s := strconv.FormatInt(n, 16) // s == "61" (hexadecimal)
	fmt.Printf("%T, %v\n",s,s)

	var nInt32 int = 100
    nstring := strconv.FormatInt(int64(nInt32), 10)
    fmt.Printf("%T, %v\n",nstring,nstring)

}