package main

import (
	"github.com/zhangdeman/go-framework/core/boot"
	"os"
	"fmt"
)

func main() {
	var env string
	cmdParam := os.Args
	if len(cmdParam) < 2 {
		env = "dev"
	} else {
		env = os.Args[1]
	}

	fmt.Println("运行环境", env)
	//启动运行服务器
	boot.RunServer(env, "/Users/didi/zhangdeman/demo/src/config")
}
