package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// GetKinesis get dynamodb service
// return *kinesis.Kinesis
func (aw *Aws) GetKinesis() *kinesis.Kinesis {
	cfg := aw.Sessions()
	kinesis := kinesis.New(session.New(), cfg)
	return kinesis
}

// GetMessagesInput ...
func (aw *Aws) GetMessagesInput() *kinesis.PutRecordInput {
	return &kinesis.PutRecordInput{}
}

// Send ...
func (aw *Aws) Send(data *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	svc := aw.GetKinesis()
	req, err := svc.PutRecord(data)
	return req, err
}

// GetShardIterator ...
func (aw *Aws) GetShardIterator() (*kinesis.GetShardIteratorOutput, error) {
	svc := aw.GetKinesis()
	dsIter := &kinesis.GetShardIteratorInput{}
	dsIter.SetStreamName(os.Getenv("KINESIS_STREAM_NAME"))
	dsIter.SetShardId(os.Getenv("KINESIS_SHARD_ID"))
	dsIter.SetShardIteratorType(os.Getenv("KINESIS_SHARD_TYPE"))
	shardIter, err := svc.GetShardIterator(dsIter)
	return shardIter, err
}

// GetRecord ...
func (aw *Aws) GetRecord(data *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
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
