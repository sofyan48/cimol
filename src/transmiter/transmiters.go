package transmiter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/provider"
)

// Transmiter ...
type Transmiter struct {
	AwsLibs  libaws.AwsInterface
	Provider provider.ProvidersInterface
}

// GetTransmiter ...
func GetTransmiter() *Transmiter {
	return &Transmiter{
		AwsLibs:  libaws.AwsHAndler(),
		Provider: provider.ProvidersHandler(),
	}
}

// ConsumerTrans ...
func (trs *Transmiter) ConsumerTrans(wg *sync.WaitGroup) {
	shardIterator, err := trs.AwsLibs.GetShardIterator()
	if err != nil {
		log.Println(err)
	}

	describeInput := trs.AwsLibs.GetDescribeInput()
	describeInput.SetStreamName("notification")
	describeInput.SetExclusiveStartShardId(os.Getenv("KINESIS_SHARD_ID"))
	for {
		err := trs.AwsLibs.WaitUntil(describeInput)
		if err != nil {
			log.Println("error Wait: ", err)
		}
		done := make(chan bool)
		go func() {
			msgInput := &kinesis.GetRecordsInput{}
			msgInput.SetShardIterator(shardIterator)

			data, err := trs.AwsLibs.Consumer(msgInput)
			if err != nil {
				log.Println(err)
			}
			itemDynamo := &entity.DynamoItem{}
			for _, i := range data.Records {
				err := json.Unmarshal([]byte(string(i.Data)), itemDynamo)
				if err != nil {
					log.Println("Error: ", err)
				}
				trs.intercepActionShard(itemDynamo)
			}
			close(done)
			shardIterator = *data.NextShardIterator
			return
		}()
		<-done
		time.Sleep(5 * time.Second)
	}

}

func (trs *Transmiter) intercepActionShard(data *entity.DynamoItem) {
	if data.History["wavecell"] != nil {
		fmt.Println("HISTORY OKE")
	} else {
		fmt.Println("HISTORY NULL")
	}
}

func infobipActionShard() {}

func wavecellActionShard() {}
