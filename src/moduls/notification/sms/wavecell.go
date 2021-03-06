package sms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// WavecellActionShard ...
func (sms *SMSModuls) WavecellActionShard(history string, payload *entity.HistoryItem) {
	reformatPayload := &entity.WavecellRequest{}
	reformatPayload.Destination = payload.Payload.Msisdn
	reformatPayload.Source = os.Getenv("WAVECELL_ACC_ID")
	reformatPayload.Text = payload.Payload.Text
	reformatPayload.ClientMessageID = payload.CallbackData
	reformatPayload.DLRCallback = os.Getenv("WAVECELL_CALLBACK_URL")
	if checkEnvironment() {
		_, err := sms.updateDynamoTransmitt(payload.CallbackData, "SENDED", "", payload)
		if err != nil {
			sms.Logs.Write("Transmitter", err.Error())
		}
		return
	}
	wavecelSendURL := "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID_GENERAL") + "/single"
	if payload.Payload.OTP == true {
		wavecelSendURL = "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID") + "/single"
	}
	wavecellReformatPayload, err := json.Marshal(reformatPayload)

	client, err := sms.Requester.CLIENT("POST", wavecelSendURL, wavecellReformatPayload)
	if err != nil {
		sms.Logs.Write("Transmitter", err.Error())
	}
	requester := &http.Client{}
	client.Header.Set("Content-Type", "application/json")
	client.Header.Set("Authorization", "Bearer "+os.Getenv("WAVECELL_ACC_TOKEN"))
	response, err := requester.Do(client)
	if err != nil {
		log.Println("Wavecell Transmitter: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		sms.Logs.Write("Transmitter", err.Error())
	}
	wavecellResponse := &entity.WavecellResponse{}
	json.Unmarshal(body, wavecellResponse)
	bodyResult := map[string]string{
		history: string(body),
	}

	bodyResultHistory, _ := json.Marshal(bodyResult)
	_, err = sms.updateDynamoTransmitt(payload.CallbackData,
		wavecellResponse.Status.Code,
		string(bodyResultHistory), payload)
	if err != nil {
		sms.Logs.Write("Transmitter", err.Error())
	}
	sms.Logs.Write("SMS SEND", string(body))
}
