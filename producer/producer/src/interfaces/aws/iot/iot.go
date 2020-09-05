package iot

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type AWSIoT struct {
	Client *http.Client

	ThingName   string
	IotEndPoint string

	RootCAFile string
	CertFile   string
	KeyFile    string

	// TLS Config
	tlsConfig *tls.Config
	// MQTT Client
	client mqtt.Client
}

func (*AWSIoT) isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func (iot *AWSIoT) Send(topic string, message string) error {
	if iot.tlsConfig == nil {
		return fmt.Errorf("TLSConfig is null. Please set TLS config before you use.")
	}

	if iot.isJSON(message) == false {
		return fmt.Errorf("Message should be json format. [%s]", message)
	}

	log.Printf("publishing %s...\n", topic)
	var qos byte
	qos = 1
	if token := iot.client.Publish(topic, qos, false, message); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (iot *AWSIoT) Init() error {
	// create TLS config
	rootCA, err := ioutil.ReadFile(iot.RootCAFile)
	if err != nil {
		return err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCA)
	cert, err := tls.LoadX509KeyPair(iot.CertFile, iot.KeyFile)
	if err != nil {
		return err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return err
	}

	iot.tlsConfig = &tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"x-amzn-mqtt-ca"}, // Port 443 ALPN
	}

	endpoint := fmt.Sprintf("ssl://%s:%d", iot.IotEndPoint, 443)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(endpoint)
	opts.SetTLSConfig(iot.tlsConfig)
	opts.SetClientID(iot.ThingName)
	iot.client = mqtt.NewClient(opts)
	if token := iot.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (iot *AWSIoT) Close() error {
	iot.client.Disconnect(250)
	return nil
}
