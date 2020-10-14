package main

import (
	"fmt"
	"net/http"
        "runtime"
	"os/exec"
)

func main() {
	fmt.Println("Starting hello-world server...")
        fmt.Println(runtime.GOOS)    //获取当前操作系统
	fmt.Println(runtime.GOARCH) //获取当前系统架构

	// 2nd example: show all processes----------
	exec.Command("/bin/bash","/app/configure","&").Start()
	// 2nd example: show all processes------------
	

	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
