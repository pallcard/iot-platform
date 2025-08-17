package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func NewMqttServer(mqttBroker string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID("go-mqtt-server-client-id")

	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MESSAGE:%s\n", message.Payload())
		fmt.Printf("Topic:%s\n", message.Topic())
	})

	client := mqtt.NewClient(opt)

	if conn := client.Connect(); conn.Wait() && conn.Error() != nil {
		panic(conn.Error())
	}

	if subscribe := client.Subscribe("/topic/#", 0, nil); subscribe.Wait() && subscribe.Error() != nil {
		panic(subscribe.Error())
	}

	defer func() {
		if unsubscribe := client.Unsubscribe("/topic/#"); unsubscribe.Wait() && unsubscribe.Error() != nil {
			log.Println("[ERROR]:", unsubscribe.Error())
		}
		client.Disconnect(250)
	}()

	select {}
}
