package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	d "example.com/sedp_server/database"
	webpush "github.com/SherClockHolmes/webpush-go"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// message callback handler
func msgCbHandler(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("TOPIC : %s\n", m.Topic())
	fmt.Printf("MSG   : %s\n", m.Payload())
	fmt.Println("Sending")
	db, err := d.InitDB("./data.db")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	var sensorID string
	var read int
	_, err = fmt.Sscanf(string(m.Payload()), "ID:%s , READ:%d", &sensorID, &read)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if read > 10 {
		pushNoti(db, fmt.Sprintf("Smoke in sensorID : %s", sensorID))
		fmt.Println("Smoke detected. The message has been sent.")
	} else {
		fmt.Println("Everything is ok.")
	}
}

func main() {
	fs := http.FileServer(http.Dir("../client"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/sub", subHandler)

	server := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID("emqx_test_client")

	opts.SetDefaultPublishHandler(msgCbHandler)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := client.Subscribe("testtopic/#", 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to : %s\n", server)
	}

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err.Error())
	}
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("Posted")
		var s d.Subscription
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(s)
		defer r.Body.Close()

		db, err := d.InitDB("./data.db")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer db.Close()
		d.AddUser(db, s)
		d.GetAllUser(db)
	}
}

func pushNoti(db *sql.DB, message string) error {

	subscriptions, err := d.GetAllUser(db)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, subscription := range subscriptions {
		s := &webpush.Subscription{
			Endpoint: subscription.Endpoint,
			Keys: webpush.Keys{
				Auth:   subscription.Keys.Auth,
				P256dh: subscription.Keys.P256dh,
			},
		}

		res, err := webpush.SendNotification([]byte(message), s, &webpush.Options{
			Subscriber:      "your-website-mail@example.com",
			VAPIDPublicKey:  "BFUDTnUGRbMR9M8JU1zz-u5irMS6Z6uRF2aJSDNYweCtxCVF76eLsgnz10ca3PTDf9AH1M7-rQ-AZhgGIkIvz2o",
			VAPIDPrivateKey: "cPFxiRP_jaKOts9WuQNIK03Ab4ewKDlyypx5Zmo0sYo",
			TTL:             30,
		})

		if err != nil {
			return err
		}

		defer res.Body.Close()

	}
	return nil
}
