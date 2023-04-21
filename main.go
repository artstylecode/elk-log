package main

import (
	"elk-log/utils"
)

func main() {
	data := make(map[string]interface{})
	data["name"] = "test"
	data["age"] = 18
	utils.Debug("123123", "test", "elk-log-test")

}
