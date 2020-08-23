package usb

import (
	"github.com/goburrow/serial"
)

type USB struct{}

func (usb *USB) Open(address string) (serial.Port, error) {
	var err error
	port, err := serial.Open(&serial.Config{Address: address})
	if err != nil {
		return nil, err
	}

	return port, nil
}
