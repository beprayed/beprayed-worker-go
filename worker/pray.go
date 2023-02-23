package worker

import (
	"encoding/json"
	"log"
	"os"
	"time"

	m "beprayed-worker-go/model"

	"github.com/nats-io/nats.go"
)

var js nats.JetStreamContext

// prayer worker is responsible to grab data from nats.io
// parse the data and insert it to the database
// main tasks are:
// 1. insert pray record into record
// 2. increment the prayer's count
func StartPrayerWorker() {
	nc, _ := nats.Connect(os.Getenv("NATS_HOST"))
	js, _ = nc.JetStream()

	sub, err := js.SubscribeSync("Pray.*", nats.Durable("pray-worker"), nats.MaxDeliver(3))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for {
		msg, err := sub.NextMsg(2 * time.Second)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Println("Received a message: ", string(msg.Data))
		var record m.PrayRecord
		if err := json.Unmarshal(msg.Data, &record); err != nil {
			log.Fatalf("Error decoding message: %v", err)
		}

		var ph m.PrayRecordModel

		if err := ph.Insert(record); err != nil {
			log.Fatalf("Error inserting pray record: %v", err)
		}

		msg.Ack()
	}
}
