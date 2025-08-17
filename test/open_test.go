package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/helper"
	"testing"
)

var openServiceAddr = "http://127.0.0.1:16001"

func TestSendMessage(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"app_key":     "app_key",
		"product_key": "1",
		"device_key":  "device_key",
		"data":        "hello!!!!",
		"sign":        "e8c68ed6bbec94d85fcfdf8bb9885cea",
	})
	rep, err := helper.HttpPost(openServiceAddr+"/sendMessage", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
