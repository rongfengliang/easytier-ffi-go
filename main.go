package main

import (
	"fmt"
	"os"
	"time"

	easytierffigo "github.com/rongfengliang/easytier-ffi-go/easytier"
)

func main() {
	data, err := os.ReadFile("app.yaml")
	if err != nil {
		panic(err)
	}
	content := string(data)
	_, cstr := easytierffigo.CString(content)
	result := easytierffigo.ParseConfig(cstr)
	fmt.Println("ParseConfig result", result)
	if result != 0 {
		fmt.Println("Error parsing config:")
	}
	result = easytierffigo.RunNetworkInstance(cstr)
	if result != 0 {
		fmt.Println("Error running network instance")
	} else {
		fmt.Println("Network instance started successfully")
	}
	for {
		fmt.Println("looping...")
		time.Sleep(1 * time.Second)
	}
}
