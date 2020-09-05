package main

import (
	"net/http"
	"os"
	"time"

	"github.com/blacknikka/kinesis-iot/interfaces/aws/iot"
)

func main() {
	iotEndPoint := os.Getenv("IOT_ENDPOINT")

	client := &http.Client{}
	awsIoT := &iot.AWSIoT{
		Client:      client,
		ThingName:   "iot",
		IotEndPoint: iotEndPoint,
		RootCAFile:  "./cert/AmazonRootCA1.pem",
		CertFile:    "./cert/iot-cert.cert.pem",
		KeyFile:     "./cert/iot-motor.private.key",
	}

	err := awsIoT.Init()
	if err != nil {
		panic(err.Error())
	}

	// // Subscribe
	// log.Printf("subscribing %s...\n", SubTopic)
	// if token := client.Subscribe(SubTopic, QoS, handleMsg); token.Wait() && token.Error() != nil {
	// 	panic(fmt.Sprintf("failed to subscribe %s: %v", SubTopic, token.Error()))
	// }

	awsIoT.Send("iot/stats", `{"message": "こんにちは"}`)

	for {
		time.Sleep(10 * time.Second)
	}
}
