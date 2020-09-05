package iot

type AWSIoT interface {
	Send(topic string, message string) error
	Close() error
	Init() error
}
