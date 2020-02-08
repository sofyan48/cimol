package transmiter

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
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
	msgInput.SetLimit(1)
	for {
		done := make(chan bool)
		go func() {
			data, err := trs.AwsLibs.Consumer(msgInput)
			if err != nil {
				log.Println(err)
			}
			for _, i := range data.Records {
				fmt.Println(string(i.Data))
			}
			close(done)
			return
		}()
		<-done
		time.Sleep(5 * time.Second)
	}

}
