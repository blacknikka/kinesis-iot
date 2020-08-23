package usecases

import (
	"github.com/blacknikka/kinesis-iot/usecases/usb"
)

type SendToUSB struct {
	Serial usb.USBInterface
}

func (s *SendToUSB) SentToUSB(message string) error {
	err := s.Serial.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}
