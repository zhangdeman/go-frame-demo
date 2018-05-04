package controller

import "fmt"
import "github.com/zhangdeman/go-framework/core/snowflake"

func TestMethod() map[string]interface{}  {
	data := make(map[string]string)
	data["name"] = "zhangdeman"
	data["age"] = "22"
	data["high"] ="180"
	returnData := make(map[string]interface{})
	returnData["errCode"] = 200
	returnData["errMsg"] = "success"
	returnData["data"] = data
	return returnData
}

func TestRootMethod() map[string]interface{}  {
	fmt.Println(snowflake.GlobalLogId)
	data := make(map[string]string)
	data["name"] = "zhangdeman"
	data["age"] = "22"
	data["high"] ="180"
	returnData := make(map[string]interface{})
	returnData["errCode"] = 200
	returnData["errMsg"] = "success"
	returnData["data"] = "我是根请求"
	return returnData
}