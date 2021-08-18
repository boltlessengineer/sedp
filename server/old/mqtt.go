package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	//"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// message callback handler
func msgCbHandler(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("TOPIC : %s\n", m.Topic())
	fmt.Printf("MSG   : %s\n", m.Payload())
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	server := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID("emqx_test_client")

	//opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(msgCbHandler)
	//opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := client.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to : %s", server)
	}

	<-c
}
