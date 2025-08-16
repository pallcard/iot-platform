package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/define"
	"iot-platform/helper"
	"testing"
)

var adminServiceAdd = "http://127.0.0.1:14001"

func TestDeviceList(t *testing.T) {
	m := define.M{
		"page": 1,
		"size": 10,
	}
	data, _ := json.Marshal(m)

	header := map[string]string{
		"token": "",
	}
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPost(adminServiceAdd+"/device/list", data, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {
	m := define.M{
		"page": 1,
		"size": 10,
	}
	data, _ := json.Marshal(m)

	header := map[string]string{
		"token": "",
	}
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPost(adminServiceAdd+"/product/list", data, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
