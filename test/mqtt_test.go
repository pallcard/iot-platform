package test

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/eclipse/paho.mqtt.golang"

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().AddBroker("tcp://192.168.101.50:1883").
		SetClientID("go-test").SetUsername("get").SetPassword("123456")

	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MESSAGE:%s\n", message.Payload())
		fmt.Printf("Topic:%s\n", message.Topic())
	})

	client := mqtt.NewClient(opt)

	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		t.Fatal(conn.Error())
	}

	//if subscribe := client.Subscribe("/topic/#", 0, func(client mqtt.Client, message mqtt.Message) {
	//	fmt.Printf("subscribe_MESSAGE:%s\n", message.Payload())
	//	fmt.Printf("subscribe_Topic:%s\n", message.Topic())
	//}); subscribe.Wait() && subscribe.Error() != nil {
	//	t.Fatal(subscribe.Error())
	//}

	if publish := client.Publish("/sys/1/1/ping", 0, false, "hello"); publish.Wait() && publish.Error() != nil {
		t.Fatal(publish.Error())
	}

	time.Sleep(time.Second * 2)

	if unsubscribe := client.Unsubscribe("/topic/#"); unsubscribe.Wait() && unsubscribe.Error() != nil {
		t.Fatal(unsubscribe.Error())
	}

	client.Disconnect(250)

}
