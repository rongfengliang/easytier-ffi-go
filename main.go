package main

import (
	"fmt"
	"os"
	"time"

	easytierffigo "github.com/rongfengliang/easytier-ffi-go/easytier"
)

var max_length = 10

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
		time.Sleep(3 * time.Second)
		// Prepare a slice to hold the results
		infos := make([]easytierffigo.KeyValuePair, max_length)
		ret := easytierffigo.CollectNetworkInfos(&infos[0], uintptr(max_length))
		if ret > 0 {
			// Process the collected network info
			for i := 0; i < int(ret); i++ {
				fmt.Printf("Network Info [%d]: %s = %s\n", i, easytierffigo.CStrToGoStr(infos[i].Key), easytierffigo.CStrToGoStr(infos[i].Value))
			}
		}
	}
}
