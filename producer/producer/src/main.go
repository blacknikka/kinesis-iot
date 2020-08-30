package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	ThingName  = "iot"
	RootCAFile = "./cert/AmazonRootCA1.pem"
	CertFile   = "./cert/iot-cert.cert.pem"
	KeyFile    = "./cert/iot-motor.private.key"
	PubTopic   = "iot/stats"
	PubMsg     = `{"message": "こんにちは"}`
	QoS        = 1
)

func main() {
	iotEndPoint := os.Getenv("IOT_ENDPOINT")

	// ブローカーに接続
	tlsConfig, err := newTLSConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to construct tls config: %v", err))
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("ssl://%s:%d", iotEndPoint, 443))
	opts.SetTLSConfig(tlsConfig)
	opts.SetClientID(ThingName)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("failed to connect broker: %v", token.Error()))
	}
	defer client.Disconnect(250)

	// // Subscribe
	// log.Printf("subscribing %s...\n", SubTopic)
	// if token := client.Subscribe(SubTopic, QoS, handleMsg); token.Wait() && token.Error() != nil {
	// 	panic(fmt.Sprintf("failed to subscribe %s: %v", SubTopic, token.Error()))
	// }

	// Publish
	log.Printf("publishing %s...\n", PubTopic)
	if token := client.Publish(PubTopic, QoS, false, PubMsg); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("failed to publish %s: %v", PubTopic, token.Error()))
	}

	for {
		time.Sleep(10 * time.Second)
	}
}

func newTLSConfig() (*tls.Config, error) {
	rootCA, err := ioutil.ReadFile(RootCAFile)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCA)
	cert, err := tls.LoadX509KeyPair(CertFile, KeyFile)
	if err != nil {
		return nil, err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
		NextProtos:         []string{"x-amzn-mqtt-ca"}, // Port 443 ALPN
	}, nil
}

func handleMsg(_ mqtt.Client, msg mqtt.Message) {
	fmt.Println(msg)
}
