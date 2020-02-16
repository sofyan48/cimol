package libfcm

import (
	"log"
	"os"

	"github.com/appleboy/go-fcm"
)

// FCMNotif ...
type FCMNotif struct {
}

// FCMNotifHandler ...
func FCMNotifHandler() *FCMNotif {
	return &FCMNotif{}
}

// FCMNotifInterface ...
type FCMNotifInterface interface {
	Send(deviceToken string, data map[string]string)
}

// Send ...
func (push *FCMNotif) Send(deviceToken string, data map[string]string) {
	msg := &fcm.Message{
		To:   deviceToken,
		Data: data,
	}

	client, err := fcm.NewClient(os.Getenv("FCM_APIKEY"))
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
}
