package main

import (
	"fmt"
	"net/http"
        "runtime"
	"os/exec"
	"io/ioutil"
)

func main() {
	fmt.Println("Starting hello-world server...")
        fmt.Println(runtime.GOOS)    //获取当前操作系统
	fmt.Println(runtime.GOARCH) //获取当前系统架构

	// 2nd example: show all processes----------
	cmd := exec.Command("/bin/bash","/app/configure", " >> file.log 2>&1" , "&" )
	err := cmd.Start()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}	
	// 2nd example: show all processes------------
	
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
	  fmt.Println(f.Name())
	}

	

	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
	fmt.Fprintln(w, "path =" + r.URL.Path)
	files, _ := ioutil.ReadDir("." + r.URL.Path)
	for _, f := range files {
	  fmt.Fprintln(w, f.Name())
	}
}
