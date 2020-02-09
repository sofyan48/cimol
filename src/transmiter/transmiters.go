package transmiter

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
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
			// msgInput.SetLimit(1)
			data, err := trs.AwsLibs.Consumer(msgInput)
			if err != nil {
				log.Println(err)
			}
			for _, i := range data.Records {
				fmt.Println("Shard Data: ", data.String())
				fmt.Println("Kinesis Data", string(i.Data))
			}
			close(done)
			shardIterator = *data.NextShardIterator
			return
		}()
		<-done
		time.Sleep(5 * time.Second)
	}

}

func (trs *Transmiter) intercepActionShard() {

}

func infobipActionShard() {}

func wavecellActionShard() {}
