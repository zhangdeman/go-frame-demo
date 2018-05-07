package main

import (
	"github.com/zhangdeman/go-framework/core/boot"
)

func main() {
	//注册配置文件路径
	//conf.LoadConfigPath("/Users/didi/zhangdeman/demo/src/config")
	//启动运行服务器
	boot.RunServer("/Users/didi/zhangdeman/demo/src/config")
}
