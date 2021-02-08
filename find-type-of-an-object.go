// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {

// 	var count int = 42
// 	var message string = "go find type"
// 	var isCheck bool = true
// 	var amount float32 = 10.2
// 	//var acounts [3]string = [3]string{"Savings", "Personal", "Joint"}

// 	// fmt.Printf("variable count=%v is of type %T \n", count, count)
// 	// fmt.Printf("variable message='%v' is of type %T \n", message, message)
// 	// fmt.Printf("variable isCheck='%v' is of type %T \n", isCheck, isCheck)
// 	// fmt.Printf("variable amount=%v is of type %T \n", amount, amount)
// 	// fmt.Printf("variable acounts=%v is of type %T \n", acounts[0], acounts[0])

// 	fmt.Printf("variable count=%v is of type %v \n", count, reflect.ValueOf(count).Kind())
// 	fmt.Printf("variable message='%v' is of type %v \n", message, reflect.ValueOf(message).Kind())
// 	fmt.Printf("variable isCheck='%v' is of type %v \n", isCheck, reflect.ValueOf(isCheck).Kind())
// 	fmt.Printf("variable amount=%v is of type %v \n", amount, reflect.ValueOf(amount).Kind())

// 	fmt.Printf("%v", reflect.TypeOf(10))

// 	fmt.Println("Using type assertions")
// 	fmt.Println(typeofObject(count))
// 	fmt.Println(typeofObject(message))
// 	fmt.Println(typeofObject(isCheck))
// 	fmt.Println(typeofObject(amount))

// 	//fmt.Printf("variable acounts=%v is of type %v \n", acounts[0], reflect.TypeOf(acounts[0]))

// 	//fmt.Println(reflect.TypeOf(intType))
// 	//fmt.Println(reflect.ValueOf(intType).Kind())

// }

// // func typeofObject(variable interface{}) string {
// // 	switch variable.(type) {
// // 	case int:
// // 		return "int"
// // 	case float64:
// // 		return "float64"
// // 	case bool:
// // 		return "boolean"
// // 	case string:
// // 		return "string"
// // 	default:
// // 		return "unknown"
// // 	}
// // }

// // func typeofObject(variable interface{}) string {
// // 	return reflect.TypeOf(variable).String()
// // }

// func typeofObject(variable interface{}) string {
// 	return reflect.ValueOf(variable).Kind().String()
// }
