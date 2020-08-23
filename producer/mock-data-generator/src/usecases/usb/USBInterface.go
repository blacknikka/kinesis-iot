package usb

type USBInterface interface {
	Write(bytes []byte) error
}
