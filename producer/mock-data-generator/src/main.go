package main

import (
	"fmt"

	"github.com/blacknikka/kinesis-iot/entities"
	"github.com/blacknikka/kinesis-iot/interfaces/usb"
	"github.com/blacknikka/kinesis-iot/usecases"
)

func main() {
	fmt.Println("hello")

	sample := entities.Record{
		Type: "type",
		Log:  "aaa\nbbb",
	}
	fmt.Println(sample)

	// usbコネクタを作る
	conn := usb.USB{}
	port, err := conn.Open("/dev/ttyUSB0")
	defer func() {
		if err := port.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	usbConnector := usb.USBSerieal{
		Port: port,
	}

	usbUsecase := &usecases.SendToUSB{
		Serial: usbConnector,
	}
	err = usbUsecase.SentToUSB("text message")
	if err != nil {
		panic(err)
	}
}
