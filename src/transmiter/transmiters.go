package transmiter

import (
	"fmt"
	"log"

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
func (trs *Transmiter) ConsumerTrans() {
	shardIterator, err := trs.AwsLibs.GetShardIterator()
	if err != nil {
		log.Println(err)
	}
	msgInput := &kinesis.GetRecordsInput{}
	msgInput.SetShardIterator(shardIterator)
	data, err := trs.AwsLibs.Consumer(msgInput)
	fmt.Println(data.Records)
}
