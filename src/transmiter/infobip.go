package transmiter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

func (trs *Transmiter) infobipActionShardOTP(history string, payload *entity.HistoryItem) {
	dest := entity.InfobipDestination{}
	dest.To = payload.Payload.Msisdn
	destination := []entity.InfobipDestination{dest}
	infobipMessages := entity.InfobipMessages{}
	infobipMessages.Destinations = destination
	infobipMessages.From = os.Getenv("INFOBIP_SENDER_ID")
	infobipMessages.Text = payload.Payload.Text
	infobipMessages.NotifyContenType = "application/json"
	infobipMessages.NotifyURL = os.Getenv("INFOBIP_CALLBACK")
	infobipMessages.CallbackData = payload.CallbackData
	infobipMessagesSlice := []entity.InfobipMessages{infobipMessages}

	infobip := &entity.InfobipRequestPayload{}
	infobip.Messages = infobipMessagesSlice
	reformatPayload, err := json.Marshal(infobip)
	if err != nil {
		log.Println("Error: ", err)
	}
	if !checkEnvironment() {
		_, err := trs.updateDynamoTransmitt(payload.CallbackData, "SENDED", " ", payload)
		if err != nil {
			log.Println("Infobip Transmitter Dynamo: ", err)
		}
		return
	}
	username := os.Getenv("INFOBIP_USERNAME")
	password := os.Getenv("INFOBIP_PASSWORD")
	client, err := trs.Requester.CLIENT("POST", os.Getenv("INFOBIP_SEND_SMS_URL"), reformatPayload)
	if err != nil {
		log.Println("Error: ", err)
	}
	requester := &http.Client{}
	client.SetBasicAuth(username, password)
	client.Header.Set("Content-Type", "application/json")
	response, err := requester.Do(client)
	if err != nil {
		log.Println("Infobip Transmitter: ", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	resultsData := string(body)
	trs.updateDynamoTransmitt(payload.CallbackData, "SENDED", resultsData, payload)
	log.Println("TRANSMITTER: ", resultsData)
}
