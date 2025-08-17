package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-platform/models"
	"log"
	"strings"
)

var topic = "/sys/#"

// topic 约定
// /sys/产品key/设备key/ping
// /sys/产品key/设备key/receive

func NewMqttServer(mqttBroker, clientId, password string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).
		SetClientID(clientId).SetUsername("get").SetPassword(password)

	opt.SetDefaultPublishHandler(publishHandler)

	client := mqtt.NewClient(opt)

	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}

	if subscribe := client.Subscribe(topic, 0, nil); subscribe.Wait() && subscribe.Error() != nil {
		panic(subscribe.Error())
	}

	defer func() {
		if unsubscribe := client.Unsubscribe(topic); unsubscribe.Wait() && unsubscribe.Error() != nil {
			log.Println("[ERROR]:", unsubscribe.Error())
		}
		client.Disconnect(250)
	}()

	select {}
}

func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE:%s\n", message.Payload())
	fmt.Printf("Topic:%s\n", message.Topic())

	topicArr := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")

	if len(topicArr) >= 4 {
		if topicArr[3] == "ping" {
			if err := models.UpdateDeviceOnlineTime(topicArr[1], topicArr[2]); err != nil {
				log.Printf("[DB ERROR]:%v\n", err)
			}

		}
	}
}
