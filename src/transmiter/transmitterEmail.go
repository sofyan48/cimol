package transmiter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/otp/src/entity/http/v1"
	"github.com/sofyan48/otp/src/util/helper/libaws"
	"github.com/sofyan48/otp/src/util/helper/provider"
	"github.com/sofyan48/otp/src/util/helper/request"
)

// TransmiterEmail ...
type TransmiterEmail struct {
	AwsLibs   libaws.AwsInterface
	Provider  provider.ProvidersInterface
	Requester request.RequesterInterface
}

// GetTransmiterEmail ...
func GetTransmiterEmail() *TransmiterEmail {
	return &TransmiterEmail{
		AwsLibs:   libaws.AwsHAndler(),
		Provider:  provider.ProvidersHandler(),
		Requester: request.RequesterHandler(),
	}
}

// ConsumerTransEmail ...
func (trs *TransmiterEmail) ConsumerTransEmail(wg *sync.WaitGroup) {
	fmt.Println("Email Consumer Exec")
	shardIterator, err := trs.AwsLibs.GetShardIterator()
	if err != nil {
		log.Println(err)
	}

	describeInput := trs.AwsLibs.GetDescribeInput()
	describeInput.SetStreamName("notification")
	describeInput.SetExclusiveStartShardId(os.Getenv("KINESIS_SHARD_EMAIL"))
	for {
		err := trs.AwsLibs.WaitUntil(describeInput)
		if err != nil {
			log.Println("error Wait: ", err)
		}
		go func() {
			msgInput := &kinesis.GetRecordsInput{}
			msgInput.SetShardIterator(shardIterator)
			data, err := trs.AwsLibs.Consumer(msgInput)
			if err != nil {
				log.Println("Kinesis : ", err)
			}
			itemDynamo := &entity.DynamoItemEmail{}
			for _, i := range data.Records {
				err := json.Unmarshal([]byte(string(i.Data)), itemDynamo)
				if err != nil {
					log.Println("Error: ", err)
				}
				fmt.Println("FROM KINESIS: ", itemDynamo)
				trs.intercepActionShardEmail(itemDynamo)
			}
			shardIterator = *data.NextShardIterator
			return
		}()
		time.Sleep(3 * time.Second)
	}

}
