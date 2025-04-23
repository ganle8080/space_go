package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// 读取指定路径下的文件，这里读取项目路径下的demo.json文件
	fileDatas, err := os.ReadFile("io_stream/demo.json")
	if err != nil {
		panic(err)
	}

	objMap := map[string]interface{}{}
	json.Unmarshal(fileDatas, &objMap)

	fmt.Printf("objMap: %v\n", objMap)
}
