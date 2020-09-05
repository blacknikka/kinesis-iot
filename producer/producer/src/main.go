package main

import (
	"net/http"
	"os"

	iotInterface "github.com/blacknikka/kinesis-iot/interfaces/aws/iot"
	iotUsecase "github.com/blacknikka/kinesis-iot/usecases/aws/iot"

	"github.com/blacknikka/kinesis-iot/usecases/serve"
)

func main() {
	iotEndPoint := os.Getenv("IOT_ENDPOINT")

	client := &http.Client{}
	awsIoT := &iotInterface.AWSIoT{
		Client:      client,
		ThingName:   "iot",
		IotEndPoint: iotEndPoint,
		RootCAFile:  "./cert/AmazonRootCA1.pem",
		CertFile:    "./cert/iot-cert.cert.pem",
		KeyFile:     "./cert/iot-motor.private.key",
	}

	iot := iotUsecase.IoTUsecase{
		IoT: awsIoT,
	}

	err := iot.InitIoT()
	if err != nil {
		panic(err.Error())
	}

	serverUsecase := serve.IoTServeUsecase{
		HttpHandler: &serve.Serve{
			PostToIoT: awsIoT.Send,
		},
	}

	if err := serverUsecase.Serve(":8000"); err != nil {
		panic(err.Error())
	}

	// if err := awsIoT.Send("iot/stats", `{"message": "こんにちは"}`); err != nil {
	// 	panic(err.Error())
	// }

	// for {
	// 	time.Sleep(10 * time.Second)
	// }
}
