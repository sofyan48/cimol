package libaws

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

// GetKinesis get dynamodb service
// return *kinesis.Kinesis
func (aw *Aws) GetKinesis() *kinesis.Kinesis {
	cfg := aw.SessionsKinesis()
	kinesis := kinesis.New(session.New(), cfg)
	return kinesis
}

// SendStart ...
func (aw *Aws) SendStart(ID string, itemDynamo *entity.DynamoItem, stack string, wg *sync.WaitGroup) {
	data, err := json.Marshal(itemDynamo)
	if err != nil {
		log.Println(err)
	}
	svc := aw.GetKinesis()
	dataSend := &kinesis.PutRecordInput{}
	dataSend.SetStreamName(os.Getenv("KINESIS_STREAM_NAME"))
	dataSend.SetPartitionKey(stack)
	dataSend.SetData(data)
	_, err = svc.PutRecord(dataSend)
	if err != nil {
		log.Println("error: ", err)
	}
	wg.Done()
	return
}

// SendMail ...
func (aw *Aws) SendMail(ID string, itemDynamo *entity.DynamoItemEmail, stack string, wg *sync.WaitGroup) {
	data, err := json.Marshal(itemDynamo)
	if err != nil {
		log.Println(err)
	}
	svc := aw.GetKinesis()
	dataSend := &kinesis.PutRecordInput{}
	dataSend.SetStreamName(os.Getenv("KINESIS_STREAM_NAME"))
	dataSend.SetPartitionKey(stack)
	dataSend.SetData(data)
	_, err = svc.PutRecord(dataSend)
	if err != nil {
		log.Println("error: ", err)
	}
	wg.Done()
	return
}

// Send ...
func (aw *Aws) Send(data []byte, stack string, wg *sync.WaitGroup) (*kinesis.PutRecordOutput, error) {
	svc := aw.GetKinesis()
	dataSend := &kinesis.PutRecordInput{}
	dataSend.SetStreamName(os.Getenv("KINESIS_STREAM_NAME"))
	dataSend.SetPartitionKey(stack)
	dataSend.SetData(data)
	req, err := svc.PutRecord(dataSend)
	if err != nil {
		log.Println(err)
	}
	wg.Done()
	return req, err
}

// GetShardIterator ...
func (aw *Aws) GetShardIterator() (string, error) {
	svc := aw.GetKinesis()
	dsIter := &kinesis.GetShardIteratorInput{}
	dsIter.SetStreamName(os.Getenv("KINESIS_STREAM_NAME"))
	dsIter.SetShardId(os.Getenv("KINESIS_SHARD_ID"))
	dsIter.SetShardIteratorType(os.Getenv("KINESIS_SHARD_TYPE"))
	shardIter, err := svc.GetShardIterator(dsIter)
	return *shardIter.ShardIterator, err
}

// Consumer ...
func (aw *Aws) Consumer(data *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	svc := aw.GetKinesis()
	records, err := svc.GetRecords(data)
	return records, err
}

// GetDescribeInput ...
func (aw *Aws) GetDescribeInput() *kinesis.DescribeStreamInput {
	return &kinesis.DescribeStreamInput{}
}

// Describe ...
func (aw *Aws) Describe(data *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	svc := aw.GetKinesis()
	descData, err := svc.DescribeStream(data)
	return descData, err
}

// WaitUntil ...
func (aw *Aws) WaitUntil(data *kinesis.DescribeStreamInput) error {
	kc := aw.GetKinesis()
	return kc.WaitUntilStreamExists(data)
}

// WaitUntilNotExist ...
func (aw *Aws) WaitUntilNotExist(data *kinesis.DescribeStreamInput) error {
	kc := aw.GetKinesis()
	return kc.WaitUntilStreamNotExists(data)
}
