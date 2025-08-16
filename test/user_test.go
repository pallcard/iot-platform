package test

import (
	"encoding/json"
	"fmt"
	"iot-platform/define"
	"iot-platform/helper"
	"testing"
)

var userServiceAdd = "http://127.0.0.1:8888"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "lk12",
		"password": "lk123",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAdd+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
