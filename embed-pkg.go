package main

import _"embed"


func main() {
	
	//go:embed "hello.txt"
    var f embed.FS
    data, _ := f.ReadFile("hello.txt")
    print(string(data))
}
