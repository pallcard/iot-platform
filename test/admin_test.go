package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/define"
	"iot-platform/helper"
	"testing"
)

var adminServiceAdd = "http://127.0.0.1:14001"

var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiSWRlbnRpdHkiOiIxIiwibmFtZSI6ImxrIiwiZXhwIjoxNzU3OTU1Mjc3fQ.HIe_TyX6k-nU_83MCNg8Ng2xPP5rY-9Iz6AYAoIgRyY"

func TestDeviceList(t *testing.T) {
	header := map[string]string{
		"token": token,
	}
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpGet(fmt.Sprintf("%s/device/list?page=%d&size=%d&name=%s",
		adminServiceAdd, 1, 10, ""), headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceCreate(t *testing.T) {
	param := define.M{
		"name":             "test产品1",
		"product_identity": "test_pro1",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPost(fmt.Sprintf("%s/device/create", adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceDelete(t *testing.T) {
	param := define.M{
		"identity": "1",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpDelete(fmt.Sprintf("%s/device/delete", adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestDeviceModify(t *testing.T) {
	param := define.M{
		"identity": "e1dd2713-fe5a-4bd0-b351-900e48050050",
		"name":     "测试产品222",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPut(fmt.Sprintf("%s/device/modify", adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductList(t *testing.T) {
	header := map[string]string{
		"token": token,
	}
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpGet(fmt.Sprintf("%s/product/list?page=%d&size=%d&name=%s",
		adminServiceAdd, 1, 10, ""), headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductCreate(t *testing.T) {
	param := define.M{
		"name": "test产品1",
		"desc": "desc",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPost(fmt.Sprintf("%s/product/create",
		adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductModify(t *testing.T) {
	param := define.M{
		"identity": "1",
		"name":     "测试产品111",
		"desc":     "测试产品描述",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpPut(fmt.Sprintf("%s/product/modify", adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProductDelete(t *testing.T) {
	param := define.M{
		"identity": "2",
	}
	header := map[string]string{
		"token": token,
	}
	paramByte, _ := json.Marshal(param)
	headerByte, _ := json.Marshal(header)
	rep, err := helper.HttpDelete(fmt.Sprintf("%s/product/delete", adminServiceAdd), paramByte, headerByte...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
