package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fileDatas, err := os.ReadFile("io_stream/demo.json")
	if err != nil {
		panic(err)
	}

	objMap := map[string]interface{}{}

	json.Unmarshal(fileDatas, &objMap)

	fmt.Printf("objMap: %v\n", objMap)
}
