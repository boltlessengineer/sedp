package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	d "example.com/sedp_server/database"
	webpush "github.com/SherClockHolmes/webpush-go"
)

func main() {
	fs := http.FileServer(http.Dir("../client"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/sub", subHandler)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Sending")
		db, err := d.InitDB("./data.db")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer db.Close()
		pushNoti(db, "hello world!")
		fmt.Println("Done")
	}()

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
