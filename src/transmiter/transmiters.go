package transmiter

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// Transmiter ...
type Transmiter struct {
	AwsLibs libaws.AwsInterface
}

// GetTransmiter ...
func GetTransmiter() *Transmiter {
	return &Transmiter{
		AwsLibs: libaws.AwsHAndler(),
	}
}

// ConsumerTrans ...
func (trs *Transmiter) ConsumerTrans(wg *sync.WaitGroup) {
	shardIterator, err := trs.AwsLibs.GetShardIterator()
	if err != nil {
		log.Println(err)
	}
	msgInput := &kinesis.GetRecordsInput{}
	msgInput.SetShardIterator(shardIterator)

	for {
		itemData := &entity.StateFullKinesis{}
		data, err := trs.AwsLibs.Consumer(msgInput)
		if err != nil {
			log.Println(err)
		}
		for _, i := range data.Records {
			err = json.Unmarshal(i.Data, itemData)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println(itemData)
		fmt.Println(itemData.Data)
		time.Sleep(5 * time.Second)
	}

}
