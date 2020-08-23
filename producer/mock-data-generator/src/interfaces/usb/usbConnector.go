package usb

import (
	"github.com/goburrow/serial"
)

type USBSerieal struct {
	Port serial.Port
}

func (usb USBSerieal) Write(bytes []byte) error {
	_, err := usb.Port.Write(bytes)
	if err != nil {
		return nil
	}

	return nil
}
